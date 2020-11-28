package database

import (
	"crypto/tls"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/pkg/errors"
	"strconv"
)

type dbDriver struct {
	db driver.Database
}

func NewDatabase(config Config) (Database, error) {
	dbConn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{config.Host + ":" + strconv.Itoa(config.Port)},
		TLSConfig: &tls.Config{},
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     dbConn,
		Authentication: driver.BasicAuthentication(config.User, config.Password)})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get a database client")
	}
	dbExists, err := client.DatabaseExists(nil, config.Database)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check if database exists")
	}
	var db driver.Database
	if !dbExists {
		db, err = client.CreateDatabase(nil, config.Database, nil)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create database")
		}
	}
	db, err = client.Database(nil, config.Database)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get database")
	}
	return &dbDriver{
		db: db,
	}, nil
}

func (d *dbDriver) Collection(name string) (Collection, error) {
	return newCollection(name, d.db)
}
