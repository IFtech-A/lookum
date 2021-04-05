package model

import "time"

type User struct {
	ID           int       `json:"id"`
	FirstName    string    `json:"first_name"`
	MiddleName   string    `json:"middle_name"`
	LastName     string    `json:"last_name"`
	Mobile       string    `json:"mobile"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	Admin        bool      `json:"is_admin"`
	Vendor       bool      `json:"is_vendor"`
	RegisteredAt time.Time `json:"registered_at"`
	LastLogin    time.Time `json:"last_login"`
	Intro        string    `json:"intro"`
	Profile      string    `json:"profile"`
}

func NewUser() *User {
	return &User{}
}
