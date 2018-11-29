// Portions copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package messageServer

import (
	"github.com/spekary/goradd/pkg/log"
	"net/http"
)

// Hub maintains the set of active clients and broadcasts messages to the clients.
var hub *WebSocketHub

// helper type to synchronize the clients map using a channel
type messageType struct {
	channel string
	message map[string]interface{}	// messages will be converted to json objects. Items in the object must be json serializable.
}

type WebSocketHub struct {
	// Registered clients, first by channel, and then by pagestate.
	// This means a form or objects on a form can subscribe to multiple channels, but more
	// then one object on the same form cannot subscribe to the same channel.
	clients map[string]map[string]*Client

	// Inbound messages from the clients.
	//Broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	send chan messageType
}

func NewWebSocketHub() *WebSocketHub {
	return &WebSocketHub{
		//Broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		send: make(chan messageType),
		clients:    make(map[string]map[string]*Client),
	}
}

func (h *WebSocketHub) run() {
	for {
		select {
		case client := <-h.register:
			log.FrameworkDebugf("Client registering for channel %s pagestate %s", client.channel, client.pagestate)
			var clientsByFormstate map[string]*Client
			var ok bool
			if clientsByFormstate,ok = h.clients[client.channel]; !ok {
				clientsByFormstate = make(map[string]*Client)
				h.clients[client.channel] = clientsByFormstate
			}
			if _,ok := clientsByFormstate[client.pagestate]; !ok {
				clientsByFormstate[client.pagestate] = client
			} else {
				// The user is registering again for a particular channel. Maybe a page refresh? Close the previous channel to prevent a memory leak
				h.unregisterClient(client.channel, client.pagestate)
				clientsByFormstate[client.pagestate] = client
			}

		case client := <-h.unregister:
			log.FrameworkDebugf("Client UNregistering for channel %s pagestate %s", client.channel, client.pagestate)
			h.unregisterClient(client.channel, client.pagestate)

		case msg := <-h.send:
			if clients, ok := h.clients[msg.channel]; ok {
				log.FrameworkDebugf("Sending to channel %s - %v", msg.channel, msg.message)
				for _, client := range clients {
					client.send <- msg.message
				}
			} else {
				log.FrameworkDebugf("Could not find channel %s", msg.channel)
			}


		/* not broadcasting currently. This might change
		case message := <-h.Broadcast:
			for client := range h.clients {
				select {
				case client.send <- message: // echo
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		*/
		}
	}
}

func (h *WebSocketHub) unregisterClient(channel string, pagestate string) {
	if clientsByFormstate, ok := h.clients[channel]; ok {
		if client, ok2 := clientsByFormstate[pagestate]; ok2 {
			close (client.send)
			delete(clientsByFormstate, pagestate)
			if len(clientsByFormstate) == 0 {
				delete(h.clients, channel)
			}
		}
	}
}

func SendMessage(channel string, message map[string]interface{}) {
	hub.send <- messageType{channel, message}
}

func MakeHub() {
	hub = NewWebSocketHub()
	go hub.run()
}

func WebsocketHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
}