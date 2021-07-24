package middleware

import (
	"context"
	"my-blog/library/auth"
	"my-blog/library/jwt"
	"net/http"

	"github.com/gogf/gf/i18n/gi18n"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Authenticate(r *ghttp.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		r.Response.WriteStatusExit(http.StatusUnauthorized, g.Map{"msg": gi18n.T(context.TODO(), "TokenInvalid")})
	}

	claims, err := jwt.ParseToken(tokenString)
	if err != nil {
		r.Response.WriteStatusExit(http.StatusUnauthorized, g.Map{"msg": gi18n.T(context.TODO(), "TokenInvalid")})
	}

	res, err := g.Model(claims.Issuer).FindOne(claims.Subject)
	if err != nil || res.IsEmpty() {
		r.Response.WriteStatusExit(http.StatusUnauthorized, g.Map{"msg": gi18n.T(context.TODO(), "TokenInvalid")})
	}

	auth.New(&res)
	r.Middleware.Next()
}
