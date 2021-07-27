package shared

import (
	"net/http"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

// SimplePares 解析请求参数
func SimplePares(r *ghttp.Request, input interface{}) {
	if err := r.Parse(input); err != nil {
		if v, ok := err.(gvalid.Error); ok {
			r.Response.WriteStatusExit(http.StatusUnprocessableEntity, g.Map{
				"msg": v.Maps(),
			})
		}

		r.Response.WriteStatusExit(http.StatusUnprocessableEntity, g.Map{
			"msg": err.Error(),
		})
	}
}
