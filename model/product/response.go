package product

type GetProductResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    int    `json:"category_id"`
	Price       int    `json:"price"`
}
