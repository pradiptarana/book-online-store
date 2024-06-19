package product

type GetProductRequest struct {
	Name      string `form:"name"`
	Category  int    `form:"category_id"`
	PageNum   int    `form:"page_num,default=0"`
	PageSize  int    `form:"page_size,default=10"`
	SortBy    string `form:"sort_by,default=10"`
	OrderType string `form:"order_type,default=ascending"`
}

type GetProductFilter struct {
	Name      string
	Category  int
	PageNum   int
	PageSize  int
	SortBy    string
	OrderType string
}
