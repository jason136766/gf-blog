package service

import (
	"context"
	"errors"
	"my-blog/app/admin/define"
	"my-blog/app/dao"
	"strings"
	"time"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/database/gdb"

	"github.com/gogf/gf/i18n/gi18n"
)

var Category = new(categoryService)

type categoryService struct{}

// Store 创建分类
func (c *categoryService) Store(ctx context.Context, input *define.CategoryStore) error {
	result, err := dao.Category.Ctx(ctx).Insert(input)
	if err != nil {
		return errors.New(gi18n.T(context.TODO(), "DatabaseError"))
	}

	row, err := result.RowsAffected()
	if err != nil {
		return errors.New(gi18n.T(context.TODO(), "DatabaseError"))
	}

	if row == 0 {
		return errors.New(gi18n.T(context.TODO(), "DatabaseError"))
	}
	return nil
}

// Categories 分类列表
func (c *categoryService) Categories(ctx context.Context) (gdb.Result, error) {
	result, err := dao.Category.Ctx(ctx).Cache(time.Hour).OrderDesc("sort").All()
	if err != nil {
		return nil, errors.New(gi18n.T(context.TODO(), "DatabaseError"))
	}

	return result, nil
}

// Delete 删除分类
func (c *categoryService) Delete(ctx context.Context, ID uint64) error {
	result, err := dao.Category.Ctx(ctx).Delete("id", ID)
	if err != nil {
		return errors.New(gi18n.T(context.TODO(), "DatabaseError"))
	}

	row, err := result.RowsAffected()
	if err != nil {
		return errors.New(gi18n.T(context.TODO(), "DatabaseError"))
	}

	if row == 0 {
		return errors.New(gi18n.Tf(context.TODO(), "NotExists", gconv.String(ID)))
	}

	return nil
}

// Update 修改分类
func (c *categoryService) Update(ctx context.Context, input *define.CategoryUpdate) error {
	_, err := dao.Category.Ctx(ctx).Where("id", input.ID).Update(g.Map{
		"category_name": input.CategoryName,
		"sort":          input.Sort,
	})
	if err != nil {
		if num := strings.Index(err.Error(), "categories_category_name_unique"); num != 0 {
			return errors.New(gi18n.Tf(context.TODO(), "Exists", input.CategoryName))
		}
		return errors.New(gi18n.T(context.TODO(), "DatabaseError"))
	}
	return nil
}
