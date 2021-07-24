package api

import (
	"my-blog/library/auth"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Category = new(categoryApi)

type categoryApi struct{}

func (c *categoryApi) Store(r *ghttp.Request) {
	r.Response.WriteJsonExit(g.Map{"user": auth.User()})
}
