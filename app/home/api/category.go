package api

import (
	"gf-blog/app/home/service"
	"net/http"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Category = categoryApi{}

type categoryApi struct{}

func (c *categoryApi) Index(r *ghttp.Request) {
	result, err := service.Category.Categories(r.Context())
	if err != nil {
		r.Response.WriteStatusExit(http.StatusInternalServerError, g.Map{
			"msg": err.Error(),
		})
	}

	r.Response.WriteJsonExit(result)
}
