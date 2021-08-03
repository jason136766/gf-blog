package router

import (
	"github.com/gogf/gf/net/ghttp"
)

func v1Routes(g *ghttp.RouterGroup) {
	g.Group("/v1", func(group *ghttp.RouterGroup) {
		//group.ALL("/hello", home.Hello)
	})
}
