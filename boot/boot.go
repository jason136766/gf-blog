package boot

import (
	_ "gf-blog/packed"

	_ "gf-blog/library/validator"
)

func init() {
	//g.SetDebug(g.Cfg().GetBool("server.Debug"))
}
