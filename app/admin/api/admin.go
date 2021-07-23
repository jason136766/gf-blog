package api

import (
	"my-blog/app/admin/define"
	"my-blog/app/admin/service"
	"my-blog/library/jwt"
	"net/http"

	"github.com/gogf/gf/util/gvalid"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Admin = new(adminApi)

type adminApi struct{}

func (a *adminApi) Register(r *ghttp.Request) {
	// 定义请求参数结构体
	var input *define.AdminInput
	// 解析请求参数
	if err := r.Parse(&input); err != nil {
		// 校验错误
		if v, ok := err.(gvalid.Error); ok {
			// 返回多个错误信息
			r.Response.WriteStatusExit(http.StatusUnprocessableEntity, g.Map{
				"msg": v.Maps(),
			})
		}
		// 校验异常的信息
		r.Response.WriteStatusExit(http.StatusUnprocessableEntity, g.Map{"msg": err.Error()})
	}
	// 调用注册逻辑
	token, err := service.Register(input)
	respondWithToken(r, token, err)
}

func (a *adminApi) Login(r *ghttp.Request) {
	var input *define.AdminInput

	if err := r.Parse(&input); err != nil {
		if v, ok := err.(gvalid.Error); ok {
			r.Response.WriteStatusExit(http.StatusUnprocessableEntity, g.Map{
				"msg": v.Maps(),
			})
		}

		r.Response.WriteStatusExit(http.StatusUnprocessableEntity, g.Map{
			"msg": err.Error(),
		})
	}

	token, err := service.Login(input)
	respondWithToken(r, token, err)
}

func respondWithToken(r *ghttp.Request, token string, err error) {
	if err != nil {
		r.Response.WriteStatusExit(http.StatusUnprocessableEntity, g.Map{"msg": err.Error()})
	}

	r.Response.WriteJsonExit(g.Map{
		"access_token": token,
		"token_type":   "Bearer",
		"expires_in":   jwt.TTl,
	})
}
