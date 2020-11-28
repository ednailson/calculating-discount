package main

type Server interface {
	Run() <-chan error
	Close()
}
