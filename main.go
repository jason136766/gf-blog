package main

import (
	_ "my-blog/boot"
	_ "my-blog/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
