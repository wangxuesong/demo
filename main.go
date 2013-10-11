// main
package main

import (
	"demo/core"
	plugins "demo/plugin"
	"fmt"
	"time"
)

func main() {
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
