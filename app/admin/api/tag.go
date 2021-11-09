package api

import (
	"gf-blog/app/admin/define"
	"gf-blog/app/admin/service"
	"gf-blog/app/shared"
	"net/http"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/net/ghttp"
)

var Tag = tagApi{}

type tagApi struct{}

func (t *tagApi) Store(r *ghttp.Request) {
	var input *define.TagStore

	shared.SimplePares(r, &input)

	if err := service.Tag.Store(r.Context(), input); err != nil {
		r.Response.WriteStatusExit(http.StatusBadRequest, g.Map{"msg": err.Error()})
	}

	r.Response.WriteJsonExit(g.Map{"msg": "标签创建成功"})
}
