package boot

import (
	_ "my-blog/packed"

	_ "my-blog/library/validator"

	"github.com/gogf/gf/frame/g"
)

func init() {
	g.SetDebug(g.Cfg().GetBool("server.Debug"))
}
