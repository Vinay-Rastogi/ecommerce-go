package models

import "time"

type Product struct {
	ID           string    `json:"id"`
	StoreID      string    `json:"store_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Brand        string    `json:"brand"`
	Category     string    `json:"category"`
	Price        float64   `json:"price"`
	Rating       float64   `json:"rating"`
	Availability bool      `json:"availability"`
	CreatedAt    time.Time `json:"created_at"`
}
