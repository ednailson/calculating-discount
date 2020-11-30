package controller

import (
	"github.com/ednailson/hash-challenge/discount-calculator/database"
	"github.com/ednailson/hash-challenge/discount-calculator/domain"
	"github.com/pkg/errors"
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
	var user domain.User
	err := c.userColl.ReadById(userId, &user)
	if err != nil && err != database.ErrNotFound {
		return nil, err
	}
	var product domain.Product
	err = c.productColl.ReadById(productId, &product)
	if err != nil {
		return nil, err
	}
	discount, discountedPrice := product.CalculateDiscount(&user)
	return &Discount{
		Percentage:   discount,
		ValueInCents: int32(discountedPrice),
	}, nil
}
