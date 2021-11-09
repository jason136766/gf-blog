package service

import (
	"context"
	"errors"
	"gf-blog/app/dao"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/i18n/gi18n"
)

var Tag = tagService{}

type tagService struct{}

// Index 标签列表
func (t *tagService) Index(ctx context.Context) (gdb.Result, error) {
	result, err := dao.Tag.Ctx(ctx).All()
	if err != nil {
		return nil, errors.New(gi18n.T(ctx, "DatabaseError"))
	}

	return result, nil
}
