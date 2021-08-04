package router

import (
	"my-blog/app/admin/api"
	"my-blog/app/middleware"

	"github.com/gogf/gf/net/ghttp"
)

func adminRoutes(g *ghttp.RouterGroup) {
	g.Group("/admin", func(group *ghttp.RouterGroup) {
		group.POST("register", api.Admin.Register)
		group.POST("login", api.Admin.Login)
		group.Middleware(middleware.Authenticate)
		group.POST("categories", api.Category.Store)
		group.GET("categories", api.Category.Index)
		group.DELETE("categories/{id}", api.Category.Delete)
		group.PATCH("categories", api.Category.Update)
		group.POST("articles", api.Article.Store)
		group.GET("articles", api.Article.Index)
		group.PATCH("articles", api.Article.Update)
		group.DELETE("articles/{id}", api.Article.Delete)
	})
}
