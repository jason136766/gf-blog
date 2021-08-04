package define

import (
	"my-blog/app/shared"
)

type ArticleStore struct {
	CategoryId uint64 `p:"category_id" v:"required|integer|min:1|exists:categories,id"`
	Title      string `p:"title" v:"required|length:3,200"`
	Content    string `p:"content" v:"required"`
}

type ArticleIndexReq struct {
	CategoryID uint64 `p:"category_id" v:"integer|min:1|exists:categories,id"`
	Title      string `p:"title" v:"length:1,200"`
	shared.PageReq
}

type ArticleUpdate struct {
	ID      uint64 ` v:"required|integer|min:1|exists:articles,id"`
	Title   string `p:"title" v:"length:3,200"`
	Content string
}
