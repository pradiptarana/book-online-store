package order

import (
	"context"
	"database/sql"
	"fmt"

	model "github.com/pradiptarana/book-online-store/model/order"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (tr *OrderRepository) GetOrderHistory(filter *model.GetOrderHistoryFilter) ([]*model.Order, error) {
	whereQuery := []string{}
	params := []any{}
	if filter.InvoiceNumber != "" {
		params = append(params, filter.InvoiceNumber)
		whereQuery = append(whereQuery, "title = ? ")
	}
	if filter.Status != "" {
		params = append(params, filter.Status)
		whereQuery = append(whereQuery, "status = ? ")
	}
	query := "SELECT * FROM transaction "
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
	fmt.Println(params...)
	stmt, err := tr.db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	rows, err := stmt.Query(params...)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()
	var result []*model.Order
	for rows.Next() {
		var each = &model.Order{}
		var err = rows.Scan(&each.Id, &each.InvoiceNumber, &each.Total, &each.UserId, &each.CreatedAt, &each.Status)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		result = append(result, each)
	}
	fmt.Println(result)
	return result, nil
}

func (tr *OrderRepository) GetOrderById(id int) (*model.Order, error) {
	fmt.Println("id", id)
	query := "SELECT * FROM transaction left join order_item on transaction.id = order_item.transaction_id where transaction.id = ?"
	stmt, err := tr.db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	params := []any{id}
	rows, err := stmt.Query(params...)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()
	result := &model.Order{}
	for rows.Next() {
		item := model.OrderDetail{}
		var err = rows.Scan(&result.Id, &result.InvoiceNumber, &result.Total, &result.UserId, &result.CreatedAt, &result.Status, &item.Id, &item.OrderId, &item.ProductId, &item.Quantity, &item.Price)
		if err != nil {
			return nil, err
		}
		result.OrderItem = append(result.OrderItem, item)
	}
	return result, nil
}

func (tr *OrderRepository) GetCurrentCart(userId int) (*model.Cart, error) {
	stmt, err := tr.db.Prepare("SELECT cart.id,cart.user_id,IFNULL(cart_item.id, 0), IFNULL(cart_item.cart_id, 0),IFNULL(cart_item.product_id, 0),IFNULL(cart_item.quantity, 0) FROM cart LEFT JOIN cart_item ON cart.id = cart_item.cart_id WHERE cart.user_id = ? AND cart_item.quantity != 0")
	if err != nil {
		fmt.Println(err.Error())
		return &model.Cart{}, err
	}
	params := []any{userId}
	rows, err := stmt.Query(params...)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()
	result := &model.Cart{}
	for rows.Next() {
		var each = model.CartItem{}
		var err = rows.Scan(&result.Id, &result.UserId, &each.Id, &each.CartId, &each.ProductId, &each.Quantity)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		result.Items = append(result.Items, each)
	}
	return result, nil
}

func (tr *OrderRepository) GetCart(id int) (*model.Cart, error) {
	stmt, err := tr.db.Prepare("SELECT * FROM cart LEFT JOIN cart_item ON cart.id = cart_item.cart_id WHERE cart.id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return &model.Cart{}, err
	}
	var us model.Cart
	if err := stmt.QueryRow(id).Scan(&us.Id, &us.UserId); err != nil {
		fmt.Println()
		return &us, fmt.Errorf("error when get cart")
	}
	return &us, nil
}

func (tr *OrderRepository) AddToCart(cart *model.Cart) error {
	// Get a Tx for making transaction requests.
	tx, err := tr.db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()
	if cart.Id == 0 {
		stmt, err := tx.Prepare("insert into cart (user_id) values (?)")
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		res, err := stmt.Exec(cart.UserId)
		if err != nil {
			return err
		}

		lastInsertId, err := res.LastInsertId()
		if err != nil {
			return err
		}
		cart.Id = int(lastInsertId)
	}
	stmt, err := tx.Prepare("insert into cart_item (cart_id, product_id, quantity) values (?, ?, ?)")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(cart)
	_, err = stmt.Exec(cart.Id, cart.Items[0].ProductId, cart.Items[0].Quantity)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (tr *OrderRepository) UpdateCart(cart *model.Cart) error {
	stmt, err := tr.db.Prepare("update cart_item set quantity = ? where cart_id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	_, err = stmt.Exec(cart.Items[0].Quantity, cart.Id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (tr *OrderRepository) CreateOrder(order *model.Order) error {
	// Get a Tx for making transaction requests.
	tx, err := tr.db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	stmt, err := tx.Prepare("insert into transaction (invoice_number, total, status, user_id, created_at) values (?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	res, err := stmt.Exec(order.InvoiceNumber, order.Total, order.Status, order.UserId, order.CreatedAt)
	if err != nil {
		return err
	}
	orderId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	for i := 0; i < len(order.OrderItem); i++ {
		stmt, err = tx.Prepare("insert into order_item (transaction_id, product_id, quantity, price) values (?, ?, ?, ?)")
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		_, err = stmt.Exec(orderId, order.OrderItem[i].ProductId, order.OrderItem[i].Quantity, order.OrderItem[i].Price)
		if err != nil {
			return err
		}
	}
	stmt, err = tx.Prepare("delete from cart_item where cart_id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	_, err = stmt.Exec(order.CartId)
	if err != nil {
		return err
	}
	return tx.Commit()
}
