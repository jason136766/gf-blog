package router

import (
	"github.com/gogf/gf/net/ghttp"
)

func v1Routes(s *ghttp.Server) {
	s.Group("/", func(group *ghttp.RouterGroup) {
		//group.ALL("/hello", home.Hello)
	})
}
