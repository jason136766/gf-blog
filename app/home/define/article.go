package define

import "gf-blog/app/shared"

type ArticleIndexReq struct {
	CategoryID uint64 `p:"category_id" v:"integer|min:1|exists:categories,id"`
	TagId      uint   `p:"tag_id" v:"integer|min:1|exists:tags,id"`
	Title      string `p:"title" v:"length:1,200"`
	shared.PageReq
}

type ArticleStore struct {
	CategoryId  uint   `p:"category_id" v:"integer|min:1|exists:categories,id"`
	TagId       uint   `p:"tag_id" v:"required-without:CategoryId|integer|min:1|exists:tags,id"`
	Title       string `p:"title" v:"required|length:3,200"`
	Content     string `p:"content" v:"required"`
	ReadMinutes int    `p:"read_minutes"`
}

type ArticleUpdate struct {
	ID          uint64 `v:"required|integer|min:1|exists:articles,id"`
	CategoryId  uint   `p:"category_id" v:"integer|min:1|exists:categories,id"`
	TagId       uint   `p:"tag_id" v:"required-without:CategoryId|integer|min:1|exists:tags,id"`
	Title       string `p:"title" v:"length:3,200"`
	Content     string `p:"content" v:"required"`
	ReadMinutes int
}
