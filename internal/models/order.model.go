package models

type OrderItemModel struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type OrderModel struct {
	ID     string           `json:"id"`
	UserID string           `json:"user_id"`
	Status string           `json:"status"`
	Items  []OrderItemModel `json:"items,omitempty"`
}
