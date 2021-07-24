package router

import (
	"my-blog/app/admin/api"
	"my-blog/app/middleware"

	"github.com/gogf/gf/net/ghttp"
)

func adminRoutes(s *ghttp.Server) {
	s.Group("admin", func(group *ghttp.RouterGroup) {
		group.POST("register", api.Admin.Register)
		group.POST("login", api.Admin.Login)
		group.Middleware(middleware.Authenticate)
		group.POST("categories", api.Category.Store)
	})
}
