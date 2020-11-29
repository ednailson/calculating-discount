package database

import "github.com/pkg/errors"

var ErrReadById = errors.New("failed to read by id")
var ErrNotFound = errors.New("data not found")
var ErrInitCollection = errors.New("failed to init collection")
var ErrConnecting = errors.New("failed to connect to database")
var ErrInitDatabase = errors.New("failed to init database")
