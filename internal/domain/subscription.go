package domain

import "time"

type Subscription struct {
	ID        string
	UserID    string
	ProductID string
	StartDate time.Time
	EndDate   *time.Time
	Status    string // active | paused | cancelled
}
