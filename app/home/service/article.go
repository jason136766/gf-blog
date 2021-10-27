package service

import (
	"context"
	"errors"
	"my-blog/app/dao"
	"my-blog/app/home/define"
	"my-blog/app/shared"

	"github.com/gogf/gf/database/gdb"

	"github.com/gogf/gf/i18n/gi18n"
)

var Article = new(articleService)

type articleService struct{}

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

	result, err := m.Page(input.Page, input.PageSize).FieldsEx("content").OrderDesc("id").All()
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

// Detail 文章详情
func (a *articleService) Detail(ctx context.Context, ID uint64) (gdb.Record, error) {
	m := dao.Article.Ctx(ctx)
	record, err := m.FindOne(ID)

	if err != nil {
		return nil, errors.New(gi18n.T(ctx, "DatabaseError"))
	}

	prevID, err := m.Where("id <", ID).Value("id")
	if err != nil {
		return nil, errors.New(gi18n.T(ctx, "DatabaseError"))
	}
	nextID, err := m.Where("id >", ID).Value("id")
	if err != nil {
		return nil, errors.New(gi18n.T(ctx, "DatabaseError"))
	}

	record["prev_id"] = prevID
	record["next_id"] = nextID

	return record, nil
}

// Store 创建文章
func (a *articleService) Store(ctx context.Context, input *define.ArticleStore) error {
	result, err := dao.Article.Ctx(ctx).Insert(input)
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
