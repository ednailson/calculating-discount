package server

type Server interface {
	Run() <-chan error
	Close()
}
