package controller

import (
	"encoding/json"
	"github.com/ednailson/hash-challenge/discount-calculator/database"
	"github.com/ednailson/hash-challenge/discount-calculator/domain"
)

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

func (c *Controller) CalculateDiscount(userId, productId string) (*ProductCommand, error) {
	userRead, err := c.userColl.ReadById(userId)
	if err != nil {
		return nil, err
	}
	var user domain.User
	err = jsonDecoder(userRead, &user)
	if err != nil {
		return nil, err
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
	return &ProductCommand{
		Id:           product.Id,
		PriceInCents: product.PriceInCents,
		Title:        product.Title,
		Description:  product.Description,
		Discount: Discount{
			Percentage:   discount,
			ValueInCents: int(discountedPrice),
		},
	}, nil
}

func jsonDecoder(from, to interface{}) error {
	body, err := json.Marshal(from)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &to)
}
