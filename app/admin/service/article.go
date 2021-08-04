package service

import (
	"context"
	"errors"
	"my-blog/app/admin/define"
	"my-blog/app/dao"
	"my-blog/app/shared"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/i18n/gi18n"
)

var Article = new(articleService)

type articleService struct{}

// Store 创建文章
func (a *articleService) Store(input *define.ArticleStore) error {
	result, err := dao.Article.Insert(input)
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

// Index 文章列表
func (a *articleService) Index(ctx context.Context, input *define.ArticleIndexReq) (*shared.PageRes, error) {
	m := dao.Article.Ctx(ctx)

	if input.CategoryID != 0 {
		m = m.Where("category_id", input.CategoryID)
	}

	if input.Title != "" {
		m = m.WhereLike("title", "%"+input.Title+"%")
	}

	count, err := m.Count()
	if err != nil {
		return nil, errors.New(gi18n.T(context.TODO(), "DatabaseError"))
	}

	result, err := m.Page(input.Page, input.PageSize).OrderDesc("id").All()
	if err != nil {
		return nil, errors.New(gi18n.T(context.TODO(), "DatabaseError"))
	}

	output := &shared.PageRes{
		Page:     input.Page,
		PageSize: input.PageSize,
		Count:    uint64(count),
		Result:   result,
	}

	return output, nil
}

// Update 修改文章
func (a *articleService) Update(ctx context.Context, input *define.ArticleUpdate) error {
	_, err := dao.Article.Ctx(ctx).Where("id", input.ID).Update(g.Map{
		"title":   input.Title,
		"content": input.Content,
	})

	if err != nil {
		return errors.New(gi18n.T(context.TODO(), "DatabaseError"))
	}

	return nil
}

// Delete 删除文章
func (a *articleService) Delete(ctx context.Context, ID uint64) error {
	result, err := dao.Article.Delete("id", ID)
	if err != nil {
		return errors.New(gi18n.T(context.TODO(), "DatabaseError"))
	}

	row, err := result.RowsAffected()
	if err != nil {
		return errors.New(gi18n.T(context.TODO(), "DatabaseError"))
	}

	if row == 0 {
		s := "ID 为 " + gconv.String(ID) + " 的文章"
		return errors.New(gi18n.Tf(context.TODO(), "NotExists", s))
	}

	return nil
}
