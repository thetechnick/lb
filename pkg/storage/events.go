/*
Copyright 2019 The LB Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package storage

import (
	"github.com/thetechnick/lb/pkg/api"
)

// EventHub dispatches resource events to registered clients
type EventHub interface {
	Broadcast(old, new api.Object)
	Register() EventClient
	Unregister(c EventClient)
	Close() error
	Run() error
}

// EventClient represents a client of the event hub
type EventClient interface {
	events() chan<- Event
	Events() <-chan Event
	Close()
}

type Event struct {
	Old, New api.Object
}

type hub struct {
	clients    map[EventClient]bool
	broadcast  chan Event
	register   chan EventClient
	unregister chan EventClient
	cancelCh   chan struct{}
	doneCh     chan error
}

// NewEventHub creates a new EventHub
func NewEventHub() EventHub {
	return &hub{
		broadcast:  make(chan Event),
		register:   make(chan EventClient),
		unregister: make(chan EventClient),
		clients:    make(map[EventClient]bool),
		cancelCh:   make(chan struct{}),
		doneCh:     make(chan error),
	}
}

func (h *hub) Broadcast(old, new api.Object) {
	h.broadcast <- Event{
		Old: old,
		New: new,
	}
}

func (h *hub) Register() EventClient {
	c := newEventClient()
	h.register <- c
	return c
}

func (h *hub) Unregister(c EventClient) {
	h.unregister <- c
}

func (h *hub) Close() error {
	close(h.cancelCh)
	return <-h.doneCh
}

func (h *hub) Run() (err error) {
	defer func() { close(h.doneCh) }()
	for {
		select {
		case <-h.cancelCh:
			for c := range h.clients {
				c.Close()
				delete(h.clients, c)
			}
			return
		case c := <-h.register:
			h.clients[c] = true
		case c := <-h.unregister:
			if _, ok := h.clients[c]; ok {
				c.Close()
				delete(h.clients, c)
			}
		case message := <-h.broadcast:
			for c := range h.clients {
				select {
				case c.events() <- message:
				default:
					c.Close()
					delete(h.clients, c)
				}
			}
		}
	}
}

type client struct {
	send chan Event
}

func newEventClient() EventClient {
	return &client{
		send: make(chan Event, 100),
	}
}

func (c *client) Close() {
	close(c.send)
}

func (c *client) events() chan<- Event {
	return c.send
}

func (c *client) Events() <-chan Event {
	return c.send
}
