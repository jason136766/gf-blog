package router

import (
	"my-blog/app/home/api"
	"my-blog/app/middleware"

	"github.com/gogf/gf/net/ghttp"
)

func v1Routes(g *ghttp.RouterGroup) {
	g.Group("/v1", func(group *ghttp.RouterGroup) {
		group.GET("categories", api.Category.Index) // 分类列表
		group.GET("articles", api.Article.Index)
		group.GET("articles/{id}", api.Article.Detail)
		group.POST("login", api.User.Login)
		group.Middleware(middleware.Authenticate)
		group.POST("articles", api.Article.Store)
	})
}
