package app

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	auth "github.com/pradiptarana/book-online-store/internal/auth"
	"github.com/pradiptarana/book-online-store/internal/cache"
	dbd "github.com/pradiptarana/book-online-store/internal/db"
	env "github.com/pradiptarana/book-online-store/internal/env"
	orderRepo "github.com/pradiptarana/book-online-store/repository/order"
	productRepo "github.com/pradiptarana/book-online-store/repository/product"
	usersRepo "github.com/pradiptarana/book-online-store/repository/user"
	orderTr "github.com/pradiptarana/book-online-store/transport/api/order"
	productTr "github.com/pradiptarana/book-online-store/transport/api/product"
	usersTr "github.com/pradiptarana/book-online-store/transport/api/user"
	orderUC "github.com/pradiptarana/book-online-store/usecase/order"
	productUC "github.com/pradiptarana/book-online-store/usecase/product"
	usersUC "github.com/pradiptarana/book-online-store/usecase/user"
)

func SetupServer() *gin.Engine {
	ctx := context.Background()
	err := env.LoadEnv()
	if err != nil {
		fmt.Println(err)
		log.Fatalf("error when load env file")
	}
	db := dbd.NewDBConnection()
	myCache := cache.New[int, []byte]()
	userRepo := usersRepo.NewUserRepository(db)
	productRepo := productRepo.NewProductRepository(db)
	orderRepo := orderRepo.NewOrderRepository(db)
	userUC := usersUC.NewUserUC(userRepo)
	productUC := productUC.NewProductUC(productRepo, *myCache)
	orderUC := orderUC.NewOrderUC(orderRepo, productUC)
	userTr := usersTr.NewUsersTransport(userUC)
	productTr := productTr.NewProductTransport(productUC)
	orderTr := orderTr.NewOrderTransport(orderUC)
	router := gin.Default()
	r := router.Group("/api/v1")
	r.POST("/signup", userTr.SignUp)
	r.POST("/login", userTr.Login)

	protected := router.Group("/api/v1")
	protected.Use(auth.JwtAuthMiddleware(ctx))
	protected.GET("/product", productTr.GetProducts)
	protected.GET("/product/:id", productTr.GetProduct)
	protected.POST("/order/cart", orderTr.AddToCart)
	protected.GET("/order/cart", orderTr.GetCurrentCart)
	protected.PUT("/order/cart/:id", orderTr.UpdateCart)
	protected.POST("/order/checkout", orderTr.Checkout)
	protected.GET("/order/history", orderTr.GetOrderHistory)
	protected.GET("/order/:id", orderTr.GetOrderById)

	return router
}
