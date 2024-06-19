package order

import (
	"errors"
	"time"

	"github.com/pradiptarana/book-online-store/model/order"
	model "github.com/pradiptarana/book-online-store/model/order"
	"github.com/pradiptarana/book-online-store/repository"
	"github.com/pradiptarana/book-online-store/usecase"
)

type OrderUC struct {
	repository.OrderRepository
	usecase.ProductUsecase
}

func NewOrderUC(repo repository.OrderRepository, productUC usecase.ProductUsecase) *OrderUC {
	return &OrderUC{repo, productUC}
}

func (uc *OrderUC) AddToCart(cart *model.Cart) error {
	if cart.UserId == 0 {
		return errors.New("user not found")
	}
	currentCart, err := uc.GetCurrentCart(cart.UserId)
	if err != nil {
		return err
	}
	if currentCart.Id != 0 {
		cart.Id = currentCart.Id
	}
	err = uc.OrderRepository.AddToCart(cart)
	if err != nil {
		return err
	}
	return nil
}

func (uc *OrderUC) GetCurrentCart(userId int) (*model.Cart, error) {
	if userId == 0 {
		return nil, errors.New("user not found")
	}
	data, err := uc.OrderRepository.GetCurrentCart(userId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (uc *OrderUC) Checkout(userId int) error {
	cart, err := uc.OrderRepository.GetCurrentCart(userId)
	if err != nil {
		return err
	}
	total := 0
	order := &model.Order{
		CartId:        cart.Id,
		Status:        "complete",
		UserId:        userId,
		InvoiceNumber: uc.generateInvoice(),
		CreatedAt:     time.Now().Format(time.RFC3339),
	}
	for i := 0; i < len(cart.Items); i++ {
		product, _ := uc.GetProduct(cart.Items[i].ProductId)
		total += cart.Items[i].Quantity * product.Price
		order.OrderItem = append(order.OrderItem, model.OrderDetail{
			ProductId: product.Id,
			Quantity:  cart.Items[i].Quantity,
			Price:     product.Price,
		})
	}
	order.Total = total
	err = uc.OrderRepository.CreateOrder(order)
	if err != nil {
		return err
	}
	return nil
}

func (uc *OrderUC) generateInvoice() string {
	return "INV/" + time.Now().Format("20060102") + "/" + "1"
}

func (uc *OrderUC) GetOrderHistory(filter *model.GetOrderHistoryFilter) ([]*order.Order, error) {
	// can  be improve using cache
	data, err := uc.OrderRepository.GetOrderHistory(filter)
	if err != nil {
		return []*model.Order{}, err
	}
	return data, nil
}

func (uc *OrderUC) GetOrderById(id int) (*order.Order, error) {
	// can  be improve using cache
	data, err := uc.OrderRepository.GetOrderById(id)
	if err != nil {
		return &order.Order{}, err
	}
	for i := 0; i < len(data.OrderItem); i++ {
		product, _ := uc.GetProduct(data.OrderItem[i].ProductId)
		data.OrderItem[i].ProductName = product.Name
	}
	return data, nil
}
