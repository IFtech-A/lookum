package model

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int       `json:"id"`
	FirstName    string    `json:"first_name"`
	MiddleName   string    `json:"middle_name"`
	LastName     string    `json:"last_name"`
	Mobile       string    `json:"mobile"`
	Email        string    `json:"email"`
	Password     string    `json:"password,omitempty"`
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

func (u *User) GeneratePasswordHash() ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) GenerateToken(signKey []byte) (string, error) {

	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(u.ID),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	return claim.SignedString(signKey)
}
