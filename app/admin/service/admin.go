package service

import (
	"context"
	"encoding/hex"
	"errors"
	"my-blog/app/admin/define"
	"my-blog/app/dao"
	"my-blog/library/jwt"

	"github.com/gogf/gf/i18n/gi18n"

	"github.com/gogf/gf/crypto/gaes"

	"github.com/gogf/gf/frame/g"
)

var appKey = g.Cfg().GetBytes("server.AppKey")
var model = dao.Admin.Table

func Register(input *define.AdminInput) (string, error) {
	admin, err := dao.Admin.One("username = ?", input.Username)
	if err != nil {
		return "", err
	} else if !admin.IsEmpty() {
		return "", errors.New(gi18n.Tf(context.TODO(), "UserExists", admin["username"]))
	}

	password, err := gaes.Encrypt([]byte(input.Password), appKey)
	if err != nil {
		return "", err
	}
	input.Password = hex.EncodeToString(password)

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

	decoded, err := hex.DecodeString(admin["password"].String())
	if err != nil {
		return "", err
	}
	adminPassword, err := gaes.Decrypt(decoded, appKey)
	if err != nil {
		return "", err
	}

	if string(adminPassword) != input.Password {
		return "", errors.New(gi18n.T(context.TODO(), "UsernameOrPasswordError"))
	}

	token, err := jwt.TokenGenerator(model, admin["id"].Int64())
	if err != nil {
		return "", errors.New(gi18n.T(context.TODO(), "TokenGenerationFailed"))
	}

	return token, nil
}
