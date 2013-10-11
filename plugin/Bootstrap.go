// Bootstrap
package plugin

import (
	core "demo/core"
	"fmt"
)

type Bootstrap struct {
	channel chan interface{}
	disp    core.Router
}

func NewBootstrap() *Bootstrap {
	return new(Bootstrap)
}

func (p *Bootstrap) Init(d core.Router) {
	//p.host = d.Channel
	p.channel = make(chan interface{})
	r := d
	r.Register("/bootstrap", p.channel)
	p.disp = d
}

func (p *Bootstrap) Run() {
	m := <-p.channel
	fmt.Printf("%d\n", m)
}
