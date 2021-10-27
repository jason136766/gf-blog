package define

import "my-blog/app/shared"

type ArticleIndexReq struct {
	CategoryID uint64 `p:"category_id" v:"integer|min:1|exists:categories,id"`
	Title      string `p:"title" v:"length:1,200"`
	shared.PageReq
}

type ArticleStore struct {
	CategoryId uint   `p:"category_id" v:"required|integer|min:1|exists:categories,id"`
	Title      string `p:"title" v:"required|length:3,200"`
	Content    string `p:"content" v:"required"`
}
