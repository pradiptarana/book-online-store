package product

type Product struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	IsActive    int    `db:"is_active"`
	Category    int    `db:"category_id"`
	Price       int    `db:"price"`
}
