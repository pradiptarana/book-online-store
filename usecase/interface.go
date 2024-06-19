package usecase

import (
	"github.com/pradiptarana/book-online-store/model"
	"github.com/pradiptarana/book-online-store/model/order"
	"github.com/pradiptarana/book-online-store/model/product"
)

type UsersUsecase interface {
	SignUp(req *model.User) error
	Login(req *model.LoginRequest) (string, error)
}

type TaskUsecase interface {
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task) error
	GetTasks(req *model.GetTasksRequest) ([]*model.Task, error)
}

type ProductUsecase interface {
	GetProducts(filter *product.GetProductFilter) ([]*product.Product, error)
	GetProduct(productId int) (*product.Product, error)
}

type OrderUsecase interface {
	GetCurrentCart(userId int) (*order.Cart, error)
	AddToCart(req *order.Cart) error
	UpdateCart(req *order.Cart) error
	Checkout(userId int) error
	GetOrderHistory(filter *order.GetOrderHistoryFilter) ([]*order.Order, error)
	GetOrderById(id int) (*order.Order, error)
}
