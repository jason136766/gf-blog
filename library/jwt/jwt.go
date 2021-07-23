package jwt

import (
	"strconv"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/golang-jwt/jwt"
)

const TTl = 3600

var secret = g.Cfg().GetBytes("server.JwtSecret")

func TokenGenerator(id int64) (tokenString string, err error) {

	claims := jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: TTl,
		Subject:   strconv.FormatInt(id, 10),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}
