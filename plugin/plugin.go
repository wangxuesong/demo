// plugin
package plugin

import (
	"github.com/wangxuesong/demo/core"
)

type Plugin interface {
	Init(d core.Router)
}
