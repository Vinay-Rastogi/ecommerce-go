package main

import (
	"log"
	"net/http"

	"ecommerce/internal/config"
	"ecommerce/internal/handlers"
	"ecommerce/internal/repositories"
	"ecommerce/internal/router"
	"ecommerce/internal/search"
	"ecommerce/internal/services"
)

func main() {
	// ---------------- DB ----------------
	db := config.ConnectDB()
	defer db.Close()

	// ---------------- USERS ----------------
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// ---------------- STORES ----------------
	storeRepo := repositories.NewStoreRepo(db)
	storeService := services.NewStoreService(storeRepo)
	storeHandler := handlers.NewStoreHandler(storeService)

	// ---------------- PRODUCTS ----------------
	productRepo := repositories.NewProductRepo(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// ---------------- ORDERS ----------------
	orderRepo := repositories.NewOrderRepo(db)
	orderService := services.NewOrderService(db, orderRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	// ---------------- SUBSCRIPTIONS ----------------
	subscriptionRepo := repositories.NewSubscriptionRepo(db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)
	subscriptionHandler := handlers.NewSubscriptionHandler(subscriptionService)

	// ---------------- PAYMENTS ----------------
	paymentRepo := repositories.NewPaymentRepo(db)
	paymentService := services.NewPaymentService(paymentRepo)
	paymentHandler := handlers.NewPaymentHandler(paymentService)

	// ---------------- ELASTICSEARCH ----------------
	esClient, err := search.NewElasticClient()
	if err != nil {
		log.Fatalf("failed to connect to Elasticsearch: %v", err)
	}

	searchService := services.NewSearchService(esClient)
	searchHandler := handlers.NewSearchHandler(searchService)

	// ---------------- ROUTER ----------------
	r := router.SetupRouter(
		userHandler,
		storeHandler,
		productHandler,
		orderHandler,
		subscriptionHandler,
		paymentHandler,
		searchHandler, // 🔍 search integrated
	)

	log.Println("Server running on :3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
