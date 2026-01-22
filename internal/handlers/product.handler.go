package handlers

import (
	"encoding/json"
	"net/http"

	"ecommerce/internal/models"
	"ecommerce/internal/services"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service}
}

// POST /stores/{store_id}/products
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	storeID := mux.Vars(r)["store_id"]

	var product models.ProductModel
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateProduct(r.Context(), storeID, &product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// GET /products/{id}
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	product, err := h.service.GetProductByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
}

// GET /stores/{store_id}/products
func (h *ProductHandler) GetProductsByStore(w http.ResponseWriter, r *http.Request) {
	storeID := mux.Vars(r)["store_id"]

	products, err := h.service.GetProductsByStore(r.Context(), storeID)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

