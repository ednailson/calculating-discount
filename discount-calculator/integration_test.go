package _test

import (
	"encoding/json"
	"fmt"
	"github.com/ednailson/hash-challenge/discount-calculator/domain"
	. "github.com/ednailson/hash-challenge/discount-calculator/helper_tests"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type productCommand struct {
	Id           string   `json:"id"`
	PriceInCents int      `json:"price_in_cents"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Discount     discount `json:"discount"`
}

type discount struct {
	Percentage   float32 `json:"percentage"`
	ValueInCents int32   `json:"value_in_cents"`
}

func TestIntegration(t *testing.T) {
	g := NewGomegaWithT(t)
	userColl := MockAndTruncateCollection(g, UserCollection)
	productColl := MockAndTruncateCollection(g, ProductCollection)
	user, err := userColl.CreateDocument(nil, fakeUser())
	g.Expect(err).ShouldNot(HaveOccurred())
	product, err := productColl.CreateDocument(nil, fakeProduct())
	g.Expect(err).ShouldNot(HaveOccurred())

	resp, err := http.Get(fmt.Sprintf("http://localhost:3333/product?user_id=%s", user.Key))

	g.Expect(err).ShouldNot(HaveOccurred())
	body, err := ioutil.ReadAll(resp.Body)
	var products []productCommand
	err = json.Unmarshal(body, &products)
	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(products).Should(ConsistOf([]productCommand{
		{
			Id:           product.Key,
			PriceInCents: 2000,
			Title:        "Notebook Gamer",
			Description:  "A great notebook",
			Discount: discount{
				Percentage:   5,
				ValueInCents: 100,
			},
		},
	}))
}

func fakeUser() domain.User {
	return domain.CreateUser("Nikola", "Tesla", time.Now().UTC().AddDate(-30, 0, 0))
}

func fakeProduct() domain.Product {
	return domain.CreateProduct(2000, "Notebook Gamer", "A great notebook")
}
