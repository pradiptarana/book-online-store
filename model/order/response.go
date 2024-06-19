package order

type CartResponse struct {
	Id    int                `json:"id"`
	Items []CartItemResponse `json:"items"`
}

type CartItemResponse struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type GetOrderHistoryResponse struct {
	Id            int    `json:"id"`
	InvoiceNumber string `json:"invoice_number"`
	Total         int    `json:"total"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at"`
}

type OrderResponse struct {
	Id            int                   `json:"id"`
	InvoiceNumber string                `json:"invoice_number"`
	Total         int                   `json:"total"`
	Status        string                `json:"status"`
	CreatedAt     string                `json:"created_at"`
	OrderItem     []OrderDetailResponse `json:"order_item"`
}

type OrderDetailResponse struct {
	ProductId   int    `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
}
