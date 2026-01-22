package main

import (
	"log"
	"net/http"

	"ecommerce/internal/config"
	"ecommerce/internal/handlers"
	"ecommerce/internal/repositories"
	"ecommerce/internal/router"
	"ecommerce/internal/services"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	// User setup
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Store setup
	storeRepo := repositories.NewStoreRepo(db)
	storeService := services.NewStoreService(storeRepo)
	storeHandler := handlers.NewStoreHandler(storeService)

	// Product setup
	productRepo := repositories.NewProductRepo(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Order setup
	orderRepo := repositories.NewOrderRepo(db)
	orderService := services.NewOrderService(db,orderRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	// Subscription setup
	subscriptionRepo := repositories.NewSubscriptionRepo(db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)
	subscriptionHandler := handlers.NewSubscriptionHandler(subscriptionService)


	// Payment setup
	paymentRepo := repositories.NewPaymentRepo(db)
paymentService := services.NewPaymentService(paymentRepo)
paymentHandler := handlers.NewPaymentHandler(paymentService)


	// Router setup
	r := router.SetupRouter(userHandler, storeHandler,productHandler,orderHandler,subscriptionHandler,paymentHandler)
	log.Println("Server running on :3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
