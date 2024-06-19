package product

import (
	"database/sql"
	"fmt"

	model "github.com/pradiptarana/book-online-store/model/product"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (tr *ProductRepository) GetProducts(filter *model.GetProductFilter) ([]*model.Product, error) {
	whereQuery := []string{
		"is_active = 1 ",
	}
	params := []any{}
	if filter.Name != "" {
		params = append(params, filter.Name)
		whereQuery = append(whereQuery, "name = ? ")
	}
	if filter.Category != 0 {
		params = append(params, filter.Category)
		whereQuery = append(whereQuery, "category_id = ? ")
	}
	query := "SELECT * FROM product"
	for k, v := range whereQuery {
		if k == 0 {
			query = query + " WHERE "
		} else {
			query = query + "AND "
		}
		query = query + v
	}
	query = query + "order by id desc limit ? offset ?"
	// params = append(params, filter.SortBy)
	// params = append(params, filter.OrderType)
	params = append(params, filter.PageSize)
	params = append(params, filter.PageNum)
	fmt.Println(query)
	stmt, err := tr.db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(params...)
	rows, err := stmt.Query(params...)
	fmt.Println(rows)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()
	var result []*model.Product
	for rows.Next() {
		var each = &model.Product{}
		var err = rows.Scan(&each.Id, &each.Name, &each.Description, &each.Category, &each.Price, &each.IsActive)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		result = append(result, each)
	}
	return result, nil
}

func (tr *ProductRepository) GetProduct(id int) (*model.Product, error) {
	rows := tr.db.QueryRow("SELECT * FROM product WHERE id = ? AND is_active = 1", id)

	var us model.Product
	if err := rows.Scan(&us.Id, &us.Name, &us.Description, &us.IsActive, &us.Price, &us.Category); err != nil {
		return &us, fmt.Errorf("error when get user")
	}
	return &us, nil
}
