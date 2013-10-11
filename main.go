// main
package main

import (
	plugins "demo/biz"
	"demo/core"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World!")

	dispatcher := core.NewDispatcher()
	go dispatcher.Run()

	event := make(chan interface{})
	dispatcher.Register(plugins.ServiceStatusKey, event)

	boot := plugins.NewBootstrap()
	boot.Init(dispatcher)
	go boot.Run()

	func() {
		var v interface{}
		for {
			v = <-event

			switch v {
			case plugins.StatusBootStart:
				fmt.Println("Boot Start")
			case plugins.StatusBootEnd:
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
