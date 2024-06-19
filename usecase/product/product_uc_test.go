package product_test

import (
	"testing"

	"github.com/pradiptarana/book-online-store/internal/cache"
	"github.com/pradiptarana/book-online-store/mocks"
	model "github.com/pradiptarana/book-online-store/model/product"

	"github.com/golang/mock/gomock"

	productUC "github.com/pradiptarana/book-online-store/usecase/product"
)

func TestGetProductsSuccess(t *testing.T) {
	filter := &model.GetProductFilter{
		PageNum:  0,
		PageSize: 10,
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProductRepo := mocks.NewMockProductRepository(mockCtrl)
	productUC := productUC.NewProductUC(mockProductRepo, *cache.New[int, []byte]())

	mockProductRepo.EXPECT().GetProducts(filter).Return([]*model.Product{
		{Id: 1,
			Name:        "Product 1",
			Description: "this is Product 1",
			IsActive:    1,
			Category:    1,
			Price:       100},
	}, nil).Times(1)

	data, err := productUC.GetProducts(filter)
	if err != nil {
		t.Fail()
	}

	if len(data) != 1 {
		t.Fail()
	}
}

func TestGetProductSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProductRepo := mocks.NewMockProductRepository(mockCtrl)
	productUC := productUC.NewProductUC(mockProductRepo, *cache.New[int, []byte]())

	mockProductRepo.EXPECT().GetProduct(1).Return(&model.Product{
		Id:          1,
		Name:        "Product 1",
		Description: "this is Product 1",
		IsActive:    1,
		Category:    1,
		Price:       100,
	}, nil).Times(1)

	data, err := productUC.GetProduct(1)
	if err != nil {
		t.Fail()
	}

	if data.Id != 1 {
		t.Fail()
	}
}
