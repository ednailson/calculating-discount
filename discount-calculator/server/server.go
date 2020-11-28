package server

import (
	"fmt"
	"github.com/ednailson/hash-challenge/discount-calculator/controller"
	"github.com/ednailson/hash-challenge/discount-calculator/server/discount"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	listener   net.Listener
	grpcServer *grpc.Server
}

func CreateServer(ctrl *controller.Controller, port int) (Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}
	discountServer := discount.CreateDiscountServer(ctrl)
	grpcServer := grpc.NewServer()
	discount.RegisterDiscountServiceServer(grpcServer, discountServer)
	return &server{
		grpcServer: grpcServer,
		listener:   listener,
	}, nil
}

func (s *server) Run() <-chan error {
	var chErr chan error
	go func() {
		if err := s.grpcServer.Serve(s.listener); err != nil {
			chErr <- err
		}
	}()
	return chErr
}

func (s *server) Close() {
	s.grpcServer.GracefulStop()
}
