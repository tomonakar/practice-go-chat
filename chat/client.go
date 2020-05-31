package main

import (
	"time"

	"github.com/gorilla/websocket"
)

// client represents a single chatting user.
type client struct {
	socket *websocket.Conn // socket is the web socket for this client.
	// send   chan []byte     // send is a channel on which messages are sent.
	send chan *message
	room *room // room is the room this client is chatting in.
}

type message struct {
	Name    string
	Message string
	Time    string
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		var msg *message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			return
		}
		t := time.Now()
		layout := "2006-01-02 15:04:05"
		msg.Time = t.Format(layout)
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			return
		}
	}
}
