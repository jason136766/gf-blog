package api

import (
	"gf-blog/app/home/define"
	"gf-blog/app/home/service"
	"gf-blog/app/shared"
	"gf-blog/library/jwt"
	"net/http"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/net/ghttp"
)

var User = userApi{}

type userApi struct{}

func (u *userApi) Login(r *ghttp.Request) {
	var input *define.UserInput

	shared.SimplePares(r, &input)

	token, err := service.User.Login(r.Context(), input)
	if err != nil {
		r.Response.WriteStatusExit(http.StatusUnprocessableEntity, g.Map{"msg": err.Error()})
	}

	r.Response.WriteJsonExit(g.Map{
		"access_token": token,
		"token_type":   "Bearer",
		"expires_in":   jwt.TTl,
	})
}
