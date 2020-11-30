package controller

import (
	"encoding/json"
	"github.com/ednailson/hash-challenge/discount-calculator/database"
	"github.com/ednailson/hash-challenge/discount-calculator/domain"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var ErrDecodingData = errors.New("failed to decode data from database")

type Controller struct {
	userColl    database.Collection
	productColl database.Collection
}

func NewController(userColl, productColl database.Collection) *Controller {
	return &Controller{
		userColl:    userColl,
		productColl: productColl,
	}
}

func (c *Controller) CalculateDiscount(userId, productId string) (*Discount, error) {
	userRead, err := c.userColl.ReadById(userId)
	if err != nil && err != database.ErrNotFound {
		return nil, err
	}
	var user *domain.User
	if userRead != nil {
		err = jsonDecoder(userRead, &user)
		if err != nil {
			return nil, err
		}
	}
	productRead, err := c.productColl.ReadById(productId)
	if err != nil {
		return nil, err
	}
	var product domain.Product
	err = jsonDecoder(productRead, &product)
	if err != nil {
		return nil, err
	}
	discount, discountedPrice := product.CalculateDiscount(user)
	return &Discount{
		Percentage:   discount,
		ValueInCents: int32(discountedPrice),
	}, nil
}

func jsonDecoder(from, to interface{}) error {
	body, err := json.Marshal(from)
	if err != nil {
		log.WithError(err).Errorf("failed to marshal data")
		return ErrDecodingData
	}
	err = json.Unmarshal(body, &to)
	if err != nil {
		log.WithError(err).Errorf("failed to unmarshal data")
		return ErrDecodingData
	}
	return nil
}
