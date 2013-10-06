// plugin
package plugin

import (
	"demo/core"
)

type Plugin interface {
	Init(d *core.Dispatcher)
}
