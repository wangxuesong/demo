// main
package main

import (
	"demo/core"
	plugins "demo/plugin"
	"fmt"
	"time"
)

type Message struct {
	key   string
	value interface{}
}

type Reactor struct {
	Channel chan Message
	events  map[string]chan interface{}
}

func (r *Reactor) Register(key string, c chan interface{}) {
	r.events[key] = c
}

func (r *Reactor) dispatch() {
	m := <-r.Channel
	if c, ok := r.events[m.key]; ok {
		c <- m.value
	}
}

func NewReactor() *Reactor {
	return &Reactor{
		Channel: make(chan Message),
		events:  make(map[string]chan interface{}),
	}
}

type Plugin interface {
	Init(r *Reactor)
}

type DemoPlugin struct {
	host    chan Message
	channel chan interface{}
	Key     string
}

func (p *DemoPlugin) do(v string) {
	m := Message{p.Key, v}
	p.host <- m
	s := <-p.channel
	fmt.Println(s)
}

func (p *DemoPlugin) Init(r *Reactor) {
	fmt.Println("DemoPlugin init")
	p.host = r.Channel
	p.channel = make(chan interface{})
	r.Register(p.Key, p.channel)
}

func main() {
	reactor := NewReactor()
	demo1 := &DemoPlugin{
		//Plugin: Plugin{reactor.channel, make(chan interface{})},
		Key: "str",
	}
	demo2 := &DemoPlugin{
		//Plugin: Plugin{reactor.channel, make(chan interface{})},
		Key: "key",
	}

	var plugin Plugin
	plugin = demo1

	//demo1.Init(reactor)
	plugin.Init(reactor)
	demo2.Init(reactor)

	go demo1.do("aaa")
	go demo2.do("bbb")
	reactor.dispatch()
	reactor.dispatch()
	fmt.Println("Hello World!")
	time.Sleep(1 * time.Second)

	dispatcher := core.NewDispatcher()
	go dispatcher.Run()
	boot := plugins.NewBootstrap()
	boot.Init(dispatcher)
	go boot.Run()
	dispatcher.Fire("/bootstrap", 1)
	time.Sleep(1 * time.Second)
	dispatcher.Stop()
	time.Sleep(1 * time.Second)
}
