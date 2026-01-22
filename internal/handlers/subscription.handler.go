package handlers

import (
	"encoding/json"
	"net/http"

	"ecommerce/internal/models"
	"ecommerce/internal/services"

	"github.com/gorilla/mux"
)

type SubscriptionHandler struct {
	service *services.SubscriptionService
}

func NewSubscriptionHandler(service *services.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{service}
}

// POST /subscriptions
func (h *SubscriptionHandler) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	var sub models.SubscriptionModel
	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateSubscription(r.Context(), &sub); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sub)
}

// GET /users/{id}/subscriptions
func (h *SubscriptionHandler) GetUserSubscriptions(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	subs, err := h.service.GetUserSubscriptions(r.Context(), userID)
	if err != nil {
		http.Error(w, "Failed to fetch subscriptions", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(subs)
}
