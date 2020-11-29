package discount

import (
	"github.com/ednailson/hash-challenge/discount-calculator/controller"
	"github.com/sirupsen/logrus"
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
	logrus.WithFields(logrus.Fields{"user_id": info.UserId, "product_id": info.ProductId}).Infof("new request to calculate discount")
	discount, err := s.ctrl.CalculateDiscount(info.UserId, info.ProductId)
	if err != nil {
		logrus.WithField("error", err).Errorf("failed to calculate discount")
		return nil, err
	}
	logrus.WithFields(logrus.Fields{"discount": discount}).Infof("discount calculated")
	return &Discount{
		Percentage:   discount.Percentage,
		ValueInCents: discount.ValueInCents,
	}, nil
}
