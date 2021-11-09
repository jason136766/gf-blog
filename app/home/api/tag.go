package api

import (
	"gf-blog/app/home/service"
	"net/http"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Tag = tagApi{}

type tagApi struct{}

func (t *tagApi) Index(r *ghttp.Request) {
	result, err := service.Tag.Index(r.Context())
	if err != nil {
		r.Response.WriteStatusExit(http.StatusBadRequest, g.Map{"msg": err.Error()})
	}

	r.Response.WriteJsonExit(result)
}
