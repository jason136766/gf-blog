package service

import (
	"context"
	"errors"
	"gf-blog/app/admin/define"
	"gf-blog/app/dao"

	"github.com/gogf/gf/i18n/gi18n"
)

var Tag = tagService{}

type tagService struct{}

func (t *tagService) Store(ctx context.Context, input *define.TagStore) error {
	result, err := dao.Tag.Ctx(ctx).Insert(input)
	if err != nil {
		return errors.New(gi18n.T(ctx, "DatabaseError"))
	}

	row, err := result.RowsAffected()
	if err != nil {
		return errors.New(gi18n.T(ctx, "DatabaseError"))
	}

	if row == 0 {
		return errors.New(gi18n.T(ctx, "DatabaseError"))
	}

	return nil
}
