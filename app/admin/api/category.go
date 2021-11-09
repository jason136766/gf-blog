package api

import (
	"gf-blog/app/admin/define"
	"gf-blog/app/admin/service"
	"gf-blog/app/shared"
	"net/http"

	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"

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

func (c *categoryApi) Store(r *ghttp.Request) {
	var input *define.CategoryStore
	// 解析请求参数
	shared.SimplePares(r, &input)
	// 调用创建类别逻辑
	if err := service.Category.Store(r.Context(), input); err != nil {
		r.Response.WriteStatusExit(http.StatusBadRequest, g.Map{
			"msg": err.Error(),
		})
	}

	r.Response.WriteJsonExit(g.Map{"msg": "分类创建成功"})
}

func (c *categoryApi) Delete(r *ghttp.Request) {
	ID := r.Get("id")
	if err := gvalid.CheckValue(r.Context(), ID, "integer", nil); err != nil {
		r.Response.WriteStatusExit(http.StatusUnprocessableEntity, g.Map{
			"msg": err.Error(),
		})
	}

	err := service.Category.Delete(r.Context(), gconv.Uint64(ID))
	if err != nil {
		r.Response.WriteStatusExit(http.StatusInternalServerError, g.Map{
			"msg": err.Error(),
		})
	}

	r.Response.WriteStatusExit(http.StatusNoContent)
}

func (c *categoryApi) Update(r *ghttp.Request) {
	var input *define.CategoryUpdate
	shared.SimplePares(r, &input)

	if err := service.Category.Update(r.Context(), input); err != nil {
		r.Response.WriteStatusExit(http.StatusInternalServerError, g.Map{
			"msg": err.Error(),
		})
	}

	r.Response.WriteJsonExit(g.Map{"msg": "修改成功"})
}
