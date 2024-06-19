package cart

type Cart struct {
	Id     int `db:"id"`
	UserId int `db:"user_id"`
	Items  []CartItem
}

type CartItem struct {
	Id         int `db:"id"`
	CartItemId int `db:"cart_item_id"`
	ProductId  int `db:"product_id"`
	Quantity   int `db:"quantity"`
}
