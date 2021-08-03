package service

import (
	"context"
	"errors"
	"my-blog/app/admin/define"
	"my-blog/app/dao"
	"my-blog/library/jwt"

	"golang.org/x/crypto/bcrypt"

	"github.com/gogf/gf/i18n/gi18n"
)

var model = dao.Admin.Table

func Register(input *define.AdminInput) (string, error) {
	admin, err := dao.Admin.One("username = ?", input.Username)
	if err != nil {
		return "", err
	} else if !admin.IsEmpty() {
		return "", errors.New(gi18n.Tf(context.TODO(), "Exists", admin["username"]))
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return "", err
	}

	input.Password = string(hashed)
	result, err := dao.Admin.Insert(input)
	if err != nil {
		return "", err
	}

	row, err := result.RowsAffected()
	if row > 0 {
		adminId, _ := result.LastInsertId()
		token, err := jwt.TokenGenerator(model, adminId)
		if err != nil {
			return "", errors.New(gi18n.T(context.TODO(), "TokenGenerationFailed"))
		}

		return token, nil
	} else {
		return "", errors.New(gi18n.T(context.TODO(), "DatabaseError"))
	}
}

func Login(input *define.AdminInput) (string, error) {
	admin, err := dao.Admin.One("username = ?", input.Username)
	if err != nil {
		return "", err
	} else if admin.IsEmpty() {
		return "", errors.New(gi18n.T(context.TODO(), "UsernameOrPasswordError"))
	}

	err = bcrypt.CompareHashAndPassword(admin["password"].Bytes(), []byte(input.Password))
	if err != nil {
		return "", errors.New(gi18n.T(context.TODO(), "UsernameOrPasswordError"))
	}

	token, err := jwt.TokenGenerator(model, admin["id"].Int64())
	if err != nil {
		return "", errors.New(gi18n.T(context.TODO(), "TokenGenerationFailed"))
	}

	return token, nil
}
