package define

type CategoryStore struct {
	CategoryName string `p:"category_name" v:"required|length:2,10|not-exists:categories,category_name"`
	Sort         int8   `p:"sort" v:"integer"`
}

type CategoryUpdate struct {
	ID           uint64 `v:"required|integer"`
	CategoryName string `p:"category_name" v:"required|length:2,10|not-exists:categories,category_name"`
	Sort         int8   `p:"sort" v:"integer"`
}
