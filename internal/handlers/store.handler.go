package handlers

import (
	"encoding/json"
	"net/http"

	"ecommerce/internal/models"
	"ecommerce/internal/services"
)

type StoreHandler struct {
	service *services.StoreService
}

func NewStoreHandler(service *services.StoreService) *StoreHandler {
	return &StoreHandler{service}
}

func (h *StoreHandler) CreateStore(w http.ResponseWriter, r *http.Request) {
	var store models.StoreModel
	if err := json.NewDecoder(r.Body).Decode(&store); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateStore(r.Context(), &store); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(store)
}

func (h *StoreHandler) GetStores(w http.ResponseWriter, r *http.Request) {
	stores, err := h.service.GetStores(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch stores", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(stores)
}
