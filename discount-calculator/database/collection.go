package database

import (
	"github.com/arangodb/go-driver"
	log "github.com/sirupsen/logrus"
)

type collection struct {
	coll driver.Collection
}

func newCollection(name string, db driver.Database) (*collection, error) {
	exist, err := db.CollectionExists(nil, name)
	if err != nil {
		return nil, ErrInitCollection
	}
	var coll driver.Collection
	if !exist {
		coll, err = db.CreateCollection(nil, name, nil)
		if err != nil {
			return nil, ErrInitCollection
		}
	}
	coll, err = db.Collection(nil, name)
	return &collection{
		coll: coll,
	}, nil
}

func (c *collection) ReadById(id string, data interface{}) error {
	_, err := c.coll.ReadDocument(nil, id, data)
	if err != nil {
		if driver.IsNotFound(err) {
			return ErrNotFound
		}
		log.Infof(err.Error())
		log.WithField("error", err).Errorf("failed to read by id")
		return ErrReadById
	}
	return nil
}
