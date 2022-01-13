package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// User schema of the user table
type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Age      int64  `json:"age"`
}
type Auth struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}
