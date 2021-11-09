package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gf-blog/app/dao"
	"gf-blog/app/home/define"
	"gf-blog/app/shared"
	"math"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/database/gdb"

	"github.com/gogf/gf/i18n/gi18n"
)

var Article = articleService{}

type articleService struct{}

// Index 文章列表
func (a *articleService) Index(ctx context.Context, input *define.ArticleIndexReq) (*shared.PageRes, error) {
	m := dao.Article.Ctx(ctx)

	if input.CategoryID != 0 {
		m = m.Where("category_id", input.CategoryID)
	}

	if input.TagId != 0 {
		m = m.Where("tag_id", input.TagId)
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
	fmt.Printf("%+v", input)
	err := g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var (
			result sql.Result
			err    error
		)

		input.ReadMinutes = a.getReadMinutes(input.Content)
		result, err = dao.Article.Ctx(ctx).Insert(input)
		if err != nil {
			return err
		}

		row, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if row == 0 {
			return err
		}

		result, err = dao.Tag.Ctx(ctx).Where("id=?", input.TagId).Increment("counter", 1)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return errors.New(gi18n.T(context.TODO(), "DatabaseError"))
	}

	return nil
}

// Update 修改文章
func (a *articleService) Update(ctx context.Context, input *define.ArticleUpdate) error {

	input.ReadMinutes = a.getReadMinutes(input.Content)

	_, err := dao.Article.Ctx(ctx).Where("id", input.ID).Update(g.Map{
		"category_id":  input.CategoryId,
		"tag_id":       input.TagId,
		"title":        input.Title,
		"content":      input.Content,
		"read_minutes": input.ReadMinutes,
	})

	if err != nil {
		return errors.New(gi18n.T(context.TODO(), "DatabaseError"))
	}

	return nil
}

func (a *articleService) getReadMinutes(content string) int {
	readTime := math.Round(float64(len(content) / 300))

	if readTime > 1 {
		return int(readTime)
	}

	return 1
}
