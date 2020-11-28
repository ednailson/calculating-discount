package database

import (
	"github.com/arangodb/go-driver"
)

type collection struct {
	coll driver.Collection
}

func newCollection(name string, db driver.Database) (*collection, error) {
	exist, err := db.CollectionExists(nil, name)
	if err != nil {
		return nil, err
	}
	var coll driver.Collection
	if !exist {
		coll, err = db.CreateCollection(nil, name, nil)
		if err != nil {
			return nil, err
		}
	}
	coll, err = db.Collection(nil, name)
	return &collection{
		coll: coll,
	}, nil
}

func (c *collection) ReadById(id string) (interface{}, error) {
	var data map[string]interface{}
	_, err := c.coll.ReadDocument(nil, id, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
