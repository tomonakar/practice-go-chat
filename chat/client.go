package main

import (
	"time"

	"github.com/gorilla/websocket"
)

// client represents a single chatting user.
type client struct {
	socket *websocket.Conn // socket is the web socket for this client.
	// send   chan []byte     // send is a channel on which messages are sent.
	send     chan *message
	room     *room // room is the room this client is chatting in.
	userData map[string]interface{}
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		var msg *message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			return
		}
		// layout := "2006-01-02 15:04:05"
		// t := time.Now()
		// msg.Time = t.Format(layout)
		msg.When = time.Now()
		msg.Name = c.userData["name"].(string)
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
