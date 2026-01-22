package models

import "time"

type SubscriptionModel struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	ProductID string     `json:"product_id"`
	StartDate time.Time  `json:"start_date"`
	EndDate   *time.Time `json:"end_date,omitempty"`
	Status    string     `json:"status"`
}

