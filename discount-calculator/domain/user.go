package domain

import (
	"github.com/ednailson/hash-challenge/discount-calculator/time_now"
	"time"
)

type User struct {
	Id          string    `json:"_key,omitempty"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

func CreateUser(firstName, lastName string, dateOfBirth time.Time) User {
	return User{
		Id:          "",
		FirstName:   firstName,
		LastName:    lastName,
		DateOfBirth: dateOfBirth,
	}
}

func (u *User) IsBirthday() bool {
	now := time_now.Now()
	return now.Day() == u.DateOfBirth.Day() && now.Month() == u.DateOfBirth.Month()
}
