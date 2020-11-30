package domain

import (
	"github.com/ednailson/hash-challenge/discount-calculator/time_now"
	"time"
)

type Product struct {
	Id           string `json:"_key,omitempty"`
	PriceInCents int    `json:"price_in_cents"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

func CreateProduct(price int, title, description string) Product {
	return Product{
		Id:           "",
		PriceInCents: price,
		Title:        title,
		Description:  description,
	}
}

func (p *Product) CalculateDiscount(user *User) (float32, float32) {
	if isBlackFriday() {
		return 10, float32(p.PriceInCents) * 0.1
	}
	if user != nil && user.IsBirthday() {
		return 5, float32(p.PriceInCents) * 0.05
	}
	return 0, 0
}

func isBlackFriday() bool {
	now := time_now.Now()
	return now.Day() == 25 && now.Month() == time.November
}
