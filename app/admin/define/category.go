package define

type CategoryInput struct {
	CategoryName string `p:"category_name" v:"required|length:2,10"`
	Sort         int8   `p:"sort" v:"integer"`
}
