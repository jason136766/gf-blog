package main

import (
	_ "gf-blog/boot"
	_ "gf-blog/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
