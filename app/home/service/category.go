package service

import (
	"context"
	"errors"
	"gf-blog/app/dao"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/i18n/gi18n"
)

var Category = categoryService{}

type categoryService struct{}

// Categories 分类列表
func (c *categoryService) Categories(ctx context.Context) (gdb.Result, error) {
	result, err := dao.Category.Ctx(ctx).OrderDesc("sort").All()
	if err != nil {
		return nil, errors.New(gi18n.T(context.TODO(), "DatabaseError"))
	}

	return result, nil
}
