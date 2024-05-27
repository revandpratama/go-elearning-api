package model

import "time"

type Score struct {
	ID        int
	KrsID     int
	Score     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
