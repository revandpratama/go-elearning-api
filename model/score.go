package model

import "time"

type Model struct {
	ID        int
	KrsID     int
	Score     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
