package repository

import (
	"github.com/pradiptarana/book-online-store/model"
	orderModel "github.com/pradiptarana/book-online-store/model/order"
	productModel "github.com/pradiptarana/book-online-store/model/product"
)

//go:generate mockgen -destination=../mocks/mock_task.go -package=mocks github.com/pradiptarana/book-online-store/repository TaskRepository
type TaskRepository interface {
	CreateTask(*model.Task) error
	GetTasks(*model.GetTasksRequest) ([]*model.Task, error)
	UpdateTask(task *model.Task) error
}

//go:generate mockgen -destination=../mocks/mock_user.go -package=mocks github.com/pradiptarana/book-online-store/repository UserRepository
type UserRepository interface {
	SignUp(us *model.User) error
	GetUser(username string) (*model.User, error)
}

//go:generate mockgen -destination=../mocks/mock_product.go -package=mocks github.com/pradiptarana/book-online-store/repository ProductRepository
type ProductRepository interface {
	GetProducts(filter *productModel.GetProductFilter) ([]*productModel.Product, error)
	GetProduct(id int) (*productModel.Product, error)
}

//go:generate mockgen -destination=../mocks/mock_order.go -package=mocks github.com/pradiptarana/book-online-store/repository OrderRepository
type OrderRepository interface {
	AddToCart(cart *orderModel.Cart) error
	GetCart(cartId int) (*orderModel.Cart, error)
	UpdateCart(cart *orderModel.Cart) error
	CreateOrder(order *orderModel.Order) error
	GetOrderHistory(filter *orderModel.GetOrderHistoryFilter) ([]*orderModel.Order, error)
	GetOrderById(id int) (*orderModel.Order, error)
	GetCurrentCart(userId int) (*orderModel.Cart, error)
}
