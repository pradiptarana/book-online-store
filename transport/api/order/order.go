package order

import (
	"net/http"
	"strconv"

	"github.com/pradiptarana/book-online-store/model/order"
	model "github.com/pradiptarana/book-online-store/model/order"
	"github.com/pradiptarana/book-online-store/usecase"

	"github.com/gin-gonic/gin"
)

type OrderTransport struct {
	usecase.OrderUsecase
}

func NewOrderTransport(uc usecase.OrderUsecase) *OrderTransport {
	return &OrderTransport{uc}
}

func (ut *OrderTransport) GetOrderHistory(c *gin.Context) {
	var req model.GetOrderHistoryRequest
	if err := c.Bind(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	data, err := ut.OrderUsecase.GetOrderHistory(&model.GetOrderHistoryFilter{
		InvoiceNumber: req.InvoiceNumber,
		Status:        req.Status,
		PageNum:       req.PageNum,
		PageSize:      req.PageSize,
		OrderType:     req.OrderType,
		SortBy:        req.SortBy,
		UserId:        int(c.GetUint("userId")),
	})
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	res := []model.GetOrderHistoryResponse{}
	for i := 0; i < len(data); i++ {
		res = append(res, model.GetOrderHistoryResponse{
			Id:            data[i].Id,
			InvoiceNumber: data[i].InvoiceNumber,
			Total:         data[i].Total,
			Status:        data[i].Status,
			CreatedAt:     data[i].CreatedAt,
		})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": res})
	return
}

func (ut *OrderTransport) GetOrderById(c *gin.Context) {
	productId := c.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	data, err := ut.OrderUsecase.GetOrderById(id)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	res := model.OrderResponse{
		InvoiceNumber: data.InvoiceNumber,
		Total:         data.Total,
		Status:        data.Status,
		CreatedAt:     data.CreatedAt,
	}
	orderItem := []model.OrderDetailResponse{}
	for j := 0; j < len(data.OrderItem); j++ {
		orderItem = append(orderItem, model.OrderDetailResponse{
			ProductId:   data.OrderItem[j].ProductId,
			ProductName: data.OrderItem[j].ProductName,
			Quantity:    data.OrderItem[j].Quantity,
			Price:       data.OrderItem[j].Price,
		})
	}
	res.OrderItem = orderItem
	c.IndentedJSON(http.StatusOK, gin.H{"data": res})
	return
}

func (ut *OrderTransport) AddToCart(c *gin.Context) {
	var req model.AddToCartRequest
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := ut.OrderUsecase.AddToCart(&order.Cart{
		UserId: int(c.GetUint("userId")),
		Items: []order.CartItem{
			{
				ProductId: req.ProductId,
				Quantity:  req.Quantity,
			},
		},
	})
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Success add to cart"})
	return
}

func (ut *OrderTransport) UpdateCart(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var req model.AddToCartRequest
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = ut.OrderUsecase.UpdateCart(&order.Cart{
		Id:     id,
		UserId: int(c.GetUint("userId")),
		Items: []order.CartItem{
			{
				ProductId: req.ProductId,
				Quantity:  req.Quantity,
			},
		},
	})
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Success update cart"})
	return
}

func (ut *OrderTransport) Checkout(c *gin.Context) {
	err := ut.OrderUsecase.Checkout(int(c.GetUint("userId")))
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Success create order"})
	return
}

func (ut *OrderTransport) GetCurrentCart(c *gin.Context) {
	data, err := ut.OrderUsecase.GetCurrentCart(int(c.GetUint("userId")))
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	res := model.CartResponse{}
	res.Id = data.Id
	for i := 0; i < len(data.Items); i++ {
		if data.Items[i].ProductId != 0 {
			res.Items = append(res.Items, model.CartItemResponse{
				ProductId: data.Items[i].ProductId,
				Quantity:  data.Items[i].Quantity,
			})
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": res})
	return
}
