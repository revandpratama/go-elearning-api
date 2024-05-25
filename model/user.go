package model

import "time"

type User struct {
	ID        string
	Role      string
	Name      string
	Username  string
	Email     string
	Phone     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// type Student struct {
// 	StudentID int
// 	User
// }

// type Teacher struct {
// 	TeacherID int
// 	User
// }
