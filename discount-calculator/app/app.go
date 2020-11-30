package app

import (
	"github.com/ednailson/hash-challenge/discount-calculator/controller"
	"github.com/ednailson/hash-challenge/discount-calculator/database"
	"github.com/ednailson/hash-challenge/discount-calculator/server"
)

type App struct {
	grpcServer server.Server
}

func LoadApp(cfg Config) (*App, error) {
	db, err := database.NewDatabase(cfg.Database.Config)
	if err != nil {
		return nil, err
	}
	userColl, err := db.Collection(cfg.Database.UserCollection)
	if err != nil {
		return nil, err
	}
	productColl, err := db.Collection(cfg.Database.ProductCollection)
	if err != nil {
		return nil, err
	}
	ctrl := controller.NewController(userColl, productColl)
	grpcServer, err := server.CreateServer(ctrl, cfg.Port)
	if err != nil {
		return nil, err
	}
	return &App{
		grpcServer: grpcServer,
	}, nil
}

func (a *App) Run() <-chan error {
	return a.grpcServer.Run()
}

func (a *App) Close() {
	a.grpcServer.Close()
}
