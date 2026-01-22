package handlers

import (
	"encoding/json"
	"net/http"
	"log"  
	"ecommerce/internal/services"
)

type SearchHandler struct {
	service *services.SearchService
}

func NewSearchHandler(service *services.SearchService) *SearchHandler {
	return &SearchHandler{service: service}
}

// GET /search/products
func (h *SearchHandler) SearchProducts(w http.ResponseWriter, r *http.Request) {
	params := services.SearchParams{
		Query:     r.URL.Query().Get("q"),
		Brand:     r.URL.Query().Get("brand"),
		Category:  r.URL.Query().Get("category"),
		InStock:   r.URL.Query().Get("in_stock"),
		MinPrice:  r.URL.Query().Get("min_price"),
		MaxPrice:  r.URL.Query().Get("max_price"),
		SortBy:    r.URL.Query().Get("sort_by"),     // price | rating
		SortOrder: r.URL.Query().Get("sort_order"), // asc | desc
		Page:      r.URL.Query().Get("page"),
		Limit:     r.URL.Query().Get("limit"),
	}

	products, err := h.service.SearchProducts(r.Context(), params)
	if err != nil {
		log.Println("SearchProducts error:", err)
		http.Error(w, "failed to search products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
