package order

type GetOrderHistoryRequest struct {
	InvoiceNumber string `form:"order_id"`
	Status        string `form:"status"`
	PageNum       int    `form:"page_num,default=0"`
	PageSize      int    `form:"page_size,default=10"`
	SortBy        string `form:"sort_by,default=10"`
	OrderType     string `form:"order_type,default=ascending"`
}

type GetOrderHistoryFilter struct {
	InvoiceNumber string
	Status        string
	PageNum       int
	PageSize      int
	SortBy        string
	OrderType     string
	UserId        int
}

type AddToCartRequest struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

// type CheckoutRequest struct {
// 	CartId int `json:"cart_id"`
// }
