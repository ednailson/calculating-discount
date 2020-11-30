package discount

import (
	"github.com/ednailson/hash-challenge/discount-calculator/controller"
	log "github.com/sirupsen/logrus"
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
	log.WithFields(log.Fields{"user_id": info.UserId, "product_id": info.ProductId}).Debugf("new request to calculate discount")
	discount, err := s.ctrl.CalculateDiscount(info.UserId, info.ProductId)
	if err != nil {
		log.WithField("error", err).Errorf("failed to calculate discount")
		return nil, err
	}
	log.WithFields(log.Fields{"discount": discount}).Debugf("discount calculated")
	return &Discount{
		Percentage:   discount.Percentage,
		ValueInCents: discount.ValueInCents,
	}, nil
}
