package discount

import (
	"github.com/ednailson/hash-challenge/discount-calculator/controller"
	"golang.org/x/net/context"
)

type Server struct {
	ctrl *controller.Controller
}

func CreateDiscountServer(ctrl *controller.Controller) *Server {
	return &Server{
		ctrl: ctrl,
	}
}

func (s *Server) CalculateDiscount(ctx context.Context, info *Info) (*Discount, error) {
	discount, err := s.ctrl.CalculateDiscount(info.UserId, info.ProductId)
	if err != nil {
		return nil, err
	}
	return &Discount{
		Percentage:   discount.Percentage,
		ValueInCents: discount.ValueInCents,
	}, nil
}
