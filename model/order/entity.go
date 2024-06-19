package order

type Cart struct {
	Id     int `db:"id"`
	UserId int `db:"user_id"`
	Items  []CartItem
}

type CartItem struct {
	Id        int `db:"id"`
	CartId    int `db:"cart_id"`
	ProductId int `db:"product_id"`
	Quantity  int `db:"quantity"`
}

type Order struct {
	Id            int    `db:"id"`
	InvoiceNumber string `db:"invoice_number"`
	CartId        int    `db:"cart_id"`
	Total         int    `db:"total"`
	Status        string `db:"status"`
	UserId        int    `db:"user_id"`
	CreatedAt     string `db:"created_at"`
	OrderItem     []OrderDetail
}

type OrderDetail struct {
	Id          int `db:"id"`
	OrderId     int `db:"order_id"`
	ProductId   int `db:"product_id"`
	Quantity    int `db:"quantity"`
	Price       int `db:"price"`
	ProductName string
}
