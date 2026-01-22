package models

type ProductModel struct {
	ID           string  `json:"id"`
	StoreID      string  `json:"store_id"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Availability bool    `json:"availability"`
}
