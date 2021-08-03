package router

import (
	"my-blog/app/middleware"

	"github.com/gogf/gf/frame/g"
)

func init() {
	s := g.Server()

	// 定义客户端接受 json 格式
	s.Use(middleware.AcceptHeader)

	group := s.Group("/api")
	// 后台路由
	adminRoutes(group)

	// 前台路由
	v1Routes(group)
}
