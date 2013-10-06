// Dispatcher
package core

import (
	"fmt"
)

type Message struct {
	Key   string
	Value interface{}
}

type Dispatcher struct {
	Channel chan Message
	events  map[string]chan interface{}
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		Channel: make(chan Message),
	}
}

func (d *Dispatcher) Register(key string, c chan interface{}) {
	d.events[key] = c
}

func (d *Dispatcher) Unregister(key string, c chan interface{}) {
	if _, ok := d.events[key]; ok {
		delete(d.events, key)
	}
}

func (d *Dispatcher) Run() {
	for {
		m := <-d.Channel
		if m.Key == ".stop" {
			fmt.Println("Bye")
			return
		}

		if c, ok := d.events[m.Key]; ok {
			c <- m.Value
		}
	}
}

func (d *Dispatcher) Stop() {
	m := Message{Key: ".stop", Value: 1}
	d.Channel <- m
}
