package middleware

import (
	"gf-blog/app"
	"net/http"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/net/ghttp"
)

func AcceptHeader(r *ghttp.Request) {
	lang := r.Header.Get("Accept-Language")
	locals := []string{"zh-CN", "en"}

	if lang == "" || !app.InArray(lang, locals) {
		g.I18n().SetLanguage("zh-CN")
	} else {
		g.I18n().SetLanguage(lang)
	}

	r.Middleware.Next()

	if r.Response.Status > http.StatusOK {
		r.Response.Header().Set("Content-Type", "application/json")
	}
}
