package api

import (
	"my-blog/app/admin/define"
	"my-blog/app/admin/service"
	"my-blog/app/shared"
	"my-blog/library/auth"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Category = new(categoryApi)

type categoryApi struct{}

func (c *categoryApi) Store(r *ghttp.Request) {
	var input *define.CategoryInput
	// 解析请求参数
	shared.SimplePares(r, input)
	// 调用创建类别逻辑
	service.Store(input)
	r.Response.WriteJsonExit(g.Map{"user": auth.User()})
}
