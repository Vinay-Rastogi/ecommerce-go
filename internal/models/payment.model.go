package models

import "time"

type PaymentModel struct {
	ID          string    `json:"id"`
	OrderID     string    `json:"order_id"`
	Amount      float64   `json:"amount"`
	Status      string    `json:"status"` // pending | success | failed
	PaymentDate time.Time `json:"payment_date"`
}
