package boot

import (
	_ "my-blog/packed"

	"github.com/gogf/gf/frame/g"
)

func init() {
	g.SetDebug(g.Cfg().GetBool("server.Debug"))
}
