package discount

import (
	"github.com/ednailson/hash-challenge/discount-calculator/controller"
	"golang.org/x/net/context"
	"log"
)

type Server struct {
	ctrl *controller.Controller
}

func CreateServer(ctrl *controller.Controller) *Server {
	return &Server{
		ctrl: ctrl,
	}
}

func (s *Server) CalculateDiscount(ctx context.Context, info *Info) (*Discount, error) {
	return &Discount{
		Percentage:   10,
		ValueInCents: 100,
	}, nil
}
