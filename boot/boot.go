package boot

import (
	_ "my-blog/packed"

	_ "my-blog/library/validator"
)

func init() {
	//g.SetDebug(g.Cfg().GetBool("server.Debug"))
}
