package helper

import "golang.org/x/crypto/bcrypt"

func VerifyPassword(hashedPassword, password string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	
	return err
}