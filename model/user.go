package model

type User struct {
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
}

type Student struct {
	StudentID int
	User
}

type Teacher struct {
	TeacherID int
	User
}
