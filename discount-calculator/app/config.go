package app

import (
	"encoding/json"
	"github.com/ednailson/hash-challenge/discount-calculator/database"
	"os"
)

type Config struct {
	LogLevel int      `json:"log_level"`
	Port     int      `json:"port"`
	Database Database `json:"database"`
}

type Database struct {
	Config            database.Config `json:"config"`
	UserCollection    string          `json:"user_collection"`
	ProductCollection string          `json:"product_collection"`
}

func NewConfigFile(filename string) error {
	err := generateConfigFile(filename, configSample())
	if err != nil {
		return err
	}
	return nil
}

func generateConfigFile(filename string, config Config) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func configSample() Config {
	return Config{
		Port: 9000,
		Database: Database{
			Config: database.Config{
				Host:     "http://arangodb.service.com.br",
				Port:     8529,
				User:     "root",
				Password: "dummyPass",
				Database: "hash-db",
			},
			UserCollection:    "user-collection",
			ProductCollection: "product-collection",
		},
	}
}
