package api

import (
	"gf-blog/app/home/define"
	"gf-blog/app/home/service"
	"gf-blog/app/shared"
	"net/http"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/util/gvalid"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/net/ghttp"
)

var Article = articleApi{}

type articleApi struct{}

func (a *articleApi) Index(r *ghttp.Request) {
	var input *define.ArticleIndexReq

	shared.SimplePares(r, &input)

	result, err := service.Article.Index(r.Context(), input)
	if err != nil {
		r.Response.WriteStatusExit(http.StatusInternalServerError, g.Map{"msg": err.Error()})
	}

	r.Response.WriteJsonExit(result)
}

func (a *articleApi) Detail(r *ghttp.Request) {
	ID := r.Get("id")

	if err := gvalid.CheckValue(r.Context(), ID, "integer", nil); err != nil {
		r.Response.WriteStatusExit(http.StatusUnprocessableEntity, g.Map{"msg": err.Error()})
	}

	article, err := service.Article.Detail(r.Context(), gconv.Uint64(ID))
	if err != nil {
		r.Response.WriteStatusExit(http.StatusInternalServerError, err.Error())
	}

	r.Response.WriteJsonExit(article)
}

func (a *articleApi) Store(r *ghttp.Request) {
	var input *define.ArticleStore

	shared.SimplePares(r, &input)

	err := service.Article.Store(r.Context(), input)
	if err != nil {
		r.Response.WriteStatusExit(http.StatusInternalServerError, g.Map{"msg": err.Error()})
	}

	r.Response.WriteJsonExit(g.Map{"msg": "文章创建成功"})
}

func (a *articleApi) Update(r *ghttp.Request) {
	var input *define.ArticleUpdate

	shared.SimplePares(r, &input)

	err := service.Article.Update(r.Context(), input)
	if err != nil {
		r.Response.WriteStatusExit(http.StatusInternalServerError, g.Map{
			"msg": err.Error(),
		})
	}

	r.Response.WriteJsonExit(g.Map{"msg": "修改成功"})
}
