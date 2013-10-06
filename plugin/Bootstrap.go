// Bootstrap
package plugin

import (
	"demo/core"
)

type Bootstrap struct {
	host    chan core.Message
	channel chan interface{}
	disp    *core.Dispatcher
}

func NewBootstrap() *Bootstrap {
	return new(Bootstrap)
}

func (p *Bootstrap) Init(d *core.Dispatcher) {
	p.host = d.Channel
	p.channel = make(chan interface{})
	d.Register("/bootstrap", p.channel)
	p.disp = d
}
