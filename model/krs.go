package model

import "time"

type KRS struct {
	ID        int
	UserID    int
	SubjectID int
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
