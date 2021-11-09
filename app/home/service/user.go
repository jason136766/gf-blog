package service

import (
	"context"
	"errors"
	"gf-blog/app/dao"
	"gf-blog/app/home/define"
	"gf-blog/library/jwt"

	"github.com/gogf/gf/i18n/gi18n"
	"golang.org/x/crypto/bcrypt"
)

var User = userService{}
var model = dao.User.Table

type userService struct{}

func (u *userService) Login(ctx context.Context, input *define.UserInput) (string, error) {
	user, err := dao.User.Ctx(ctx).One("username", input.Username)
	if err != nil {
		return "", err
	} else if user.IsEmpty() {
		return "", errors.New(gi18n.T(ctx, "UsernameOrPasswordError"))
	}

	err = bcrypt.CompareHashAndPassword(user["password"].Bytes(), []byte(input.Password))
	if err != nil {
		return "", errors.New(gi18n.T(ctx, "UsernameOrPasswordError"))
	}

	token, err := jwt.TokenGenerator(model, user["id"].Int64())
	if err != nil {
		return "", errors.New(gi18n.T(ctx, "TokenGenerationFailed"))
	}

	return token, nil
}
