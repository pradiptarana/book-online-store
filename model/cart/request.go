package cart

type AddToCartRequest struct {
	ProductId int `form:"product_id"`
	Quantity  int `form:"quantity"`
}

type CheckoutRequest struct {
	CartId int `form:"cart_id"`
}
