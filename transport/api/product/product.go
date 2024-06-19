package product

import (
	"fmt"
	"net/http"
	"strconv"

	model "github.com/pradiptarana/book-online-store/model/product"
	"github.com/pradiptarana/book-online-store/usecase"

	"github.com/gin-gonic/gin"
)

type ProductTransport struct {
	usecase.ProductUsecase
}

func NewProductTransport(uc usecase.ProductUsecase) *ProductTransport {
	return &ProductTransport{uc}
}

func (ut *ProductTransport) GetProducts(c *gin.Context) {
	var req model.GetProductRequest
	if err := c.Bind(&req); err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	data, err := ut.ProductUsecase.GetProducts(&model.GetProductFilter{
		Name:      req.Name,
		Category:  req.Category,
		PageNum:   req.PageNum,
		PageSize:  req.PageSize,
		SortBy:    req.SortBy,
		OrderType: req.OrderType,
	})
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	res := []model.GetProductResponse{}
	for i := 0; i < len(data); i++ {
		res = append(res, model.GetProductResponse{
			Id:          data[i].Id,
			Name:        data[i].Name,
			Description: data[i].Description,
			Category:    data[i].Category,
			Price:       data[i].Price,
		})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": res})
	return
}

func (ut *ProductTransport) GetProduct(c *gin.Context) {
	productId := c.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	data, err := ut.ProductUsecase.GetProduct(id)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"data": model.GetProductResponse{
		Id:          data.Id,
		Name:        data.Name,
		Description: data.Description,
		Category:    data.Category,
		Price:       data.Price,
	}})
	return
}
