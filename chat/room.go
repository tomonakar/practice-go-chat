package main

import (
	"chat/trace"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/objx"
)

type room struct {
	forward chan *message    // forward is a channel that holds incoming messages that should be forwarded to the other clients.
	join    chan *client     // join is a channel for clients wishing to join the room.
	leave   chan *client     // leave is a channel for clients wishing to leave the room.
	clients map[*client]bool // clients holds all current clients in this room.
	tracer  trace.Tracer     // tracer will receive trace information of activity in the room.
	avatar  Avatar
}

// newRoom makes a new room that is ready to go.
func newRoom(avatar Avatar) *room {
	t := time.Now()
	layout := "2006-01-02 15:04:05"
	fmt.Println("Created chat room: ", t.Format(layout))

	return &room{
		forward: make(chan *message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		tracer:  trace.Off(),
		avatar:  avatar,
	}
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			// joining
			r.clients[client] = true
			r.tracer.Trace("New client joined.\nNow ", len(r.clients), " person is used.")
		case client := <-r.leave:
			// leaving
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("Client left. \nNow ", len(r.clients), " person is used.")
		case msg := <-r.forward:
			r.tracer.Trace("Message received: ", msg.Message)
			// forward message to all clients
			for client := range r.clients {
				client.send <- msg
				r.tracer.Trace(" -- sent to client")
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	authCookie, err := req.Cookie("auth")
	if err != nil {
		log.Fatal("クッキーの取得に失敗しました：", err)
		return
	}
	client := &client{
		socket:   socket,
		send:     make(chan *message, messageBufferSize),
		room:     r,
		userData: objx.MustFromBase64(authCookie.Value),
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
