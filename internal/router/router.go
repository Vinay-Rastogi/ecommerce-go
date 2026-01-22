package router

import (
	"ecommerce/internal/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter(
	userHandler *handlers.UserHandler,
	storeHandler *handlers.StoreHandler,
	productHandler *handlers.ProductHandler,
	orderHandler *handlers.OrderHandler,
	subscriptionHandler *handlers.SubscriptionHandler,
	paymentHandler *handlers.PaymentHandler,
) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")

	r.HandleFunc("/stores", storeHandler.CreateStore).Methods("POST")
	r.HandleFunc("/stores", storeHandler.GetStores).Methods("GET")

	r.HandleFunc("/stores/{store_id}/products",productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/stores/{store_id}/products",productHandler.GetProductsByStore).Methods("GET")
	r.HandleFunc("/products/{id}",productHandler.GetProductByID).Methods("GET")

	r.HandleFunc("/orders",orderHandler.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{id}",orderHandler.GetOrder).Methods("GET")
	r.HandleFunc("/users/{id}/orders",orderHandler.GetOrdersByUser).Methods("GET")

    r.HandleFunc("/subscriptions", subscriptionHandler.CreateSubscription).Methods("POST")
    r.HandleFunc("/users/{id}/subscriptions", subscriptionHandler.GetUserSubscriptions).Methods("GET")


	// Payments
    r.HandleFunc("/payments", paymentHandler.CreatePayment).Methods("POST")


	return r
}
