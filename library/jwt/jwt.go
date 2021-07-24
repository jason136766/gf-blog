package jwt

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/golang-jwt/jwt"
)

const TTl = 3600

var secret = g.Cfg().GetBytes("server.JwtSecret")

func TokenGenerator(issue string, id int64) (tokenString string, err error) {
	claims := jwt.StandardClaims{
		Issuer:    issue,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(TTl * time.Second).Unix(),
		Subject:   strconv.FormatInt(id, 10),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

func ParseToken(tokenString string) (*jwt.StandardClaims, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	claims := jwt.StandardClaims{}

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		if v, ok := err.(*jwt.ValidationError); ok {
			if v.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token 时间已过期")
			}
		}
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token 解析错误")
}
