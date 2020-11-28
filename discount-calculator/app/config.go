package app

import "github.com/ednailson/hash-challenge/discount-calculator/database"

type Config struct {
	Port     int      `json:"port"`
	Database Database `json:"database"`
}

type Database struct {
	Config            database.Config `json:"config"`
	UserCollection    string          `json:"user_collection"`
	ProductCollection string          `json:"product_collection"`
}
