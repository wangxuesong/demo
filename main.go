// main
package main

import (
	"demo/biz"
	"demo/core"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World!")

	dispatcher := core.NewDispatcher()
	go dispatcher.Run()

	event := make(chan interface{})
	dispatcher.Register(biz.ServiceStatusKey, event)

	boot := biz.NewBootstrap()
	boot.Init(dispatcher)
	go boot.Run()

	func() {
		var v interface{}
		for {
			v = <-event

			switch v {
			case biz.StatusBootStart:
				fmt.Println("Boot Start")
			case biz.StatusBootEnd:
				fmt.Println("Boot End")
				return
			default:
				fmt.Println("Unknown Status")
				return
			}
		}
	}()

	dispatcher.Stop()
	time.Sleep(1 * time.Second)
}
