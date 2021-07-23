package service

import (
	"encoding/hex"
	"errors"
	"my-blog/app/admin/define"
	"my-blog/app/dao"
	"my-blog/library/jwt"

	"github.com/gogf/gf/crypto/gaes"

	"github.com/gogf/gf/frame/g"
)

var appKey = g.Cfg().GetBytes("server.AppKey")

func Register(input *define.AdminInput) (string, error) {
	admin, err := dao.Admin.One("username = ?", input.Username)
	if err != nil {
		return "", err
	} else if admin != nil {
		err = errors.New("管理员已存在")
		return "", err
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
	nums, err := result.RowsAffected()
	if err != nil {
		err = errors.New("数据库操作异常")
	} else if nums > 0 {
		adminId, _ := result.LastInsertId()
		token, err := jwt.TokenGenerator(adminId)
		if err != nil {
			err = errors.New("token 生成失败")
			return "", err
		}

		return token, nil
	}

	return "", err
}

func Login(input *define.AdminInput) (string, error) {
	admin, err := dao.Admin.One("username = ?", input.Username)
	if err != nil {
		return "", err
	} else if admin == nil {
		err = errors.New("用户名或密码错误")
		return "", err
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
		err = errors.New("用户名或密码错误")
		return "", err
	}

	token, err := jwt.TokenGenerator(admin["id"].Int64())
	if err != nil {
		err = errors.New("token 生成失败")
		return "", err
	}

	return token, nil
}
