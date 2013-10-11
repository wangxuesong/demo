// Bootstrap
package biz

import (
	core "demo/core"
	"time"
)

type Bootstrap struct {
	channel chan interface{}
	disp    core.Router
}

func NewBootstrap() *Bootstrap {
	return new(Bootstrap)
}

func (p *Bootstrap) Init(d core.Router) {
	p.channel = make(chan interface{})
	p.disp = d
}

func (p *Bootstrap) Run() {
	p.disp.Fire(ServiceStatusKey, StatusBootStart)
	time.Sleep(1 * time.Second)
	p.disp.Fire(ServiceStatusKey, StatusBootEnd)
}
