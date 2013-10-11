// Dispatcher
package core

import (
	"fmt"
)

type Router interface {
	Register(string, chan interface{})
	Unregister(key string, c chan interface{})
	Fire(key string, value interface{})
}

type message struct {
	Key   string
	Value interface{}
}

type Dispatcher struct {
	Channel chan message
	events  map[string]chan interface{}
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		Channel: make(chan message),
		events:  make(map[string]chan interface{}),
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

func (d *Dispatcher) Fire(key string, value interface{}) {
	m := message{key, value}
	d.Channel <- m
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
	m := message{Key: ".stop", Value: 1}
	d.Channel <- m
}
