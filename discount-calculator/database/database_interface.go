package database

type Database interface {
	Collection(name string) (Collection, error)
}

type Collection interface {
	ReadById(id string, data interface{}) error
}
