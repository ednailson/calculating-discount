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

// To run this test please run the bash init.sh before
func TestIntegration(t *testing.T) {
	g := NewGomegaWithT(t)
	userColl := MockCollection(g, UserCollection)
	createdUser, err := userColl.CreateDocument(nil, fakeUser())
	g.Expect(err).ShouldNot(HaveOccurred())
	defer userColl.RemoveDocument(nil, createdUser.Key)
	resp, err := http.Get(fmt.Sprintf("http://localhost:3333/product?user_id=%s", createdUser.Key))

	g.Expect(err).ShouldNot(HaveOccurred())
	body, err := ioutil.ReadAll(resp.Body)
	var products []productCommand
	err = json.Unmarshal(body, &products)
	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(products).Should(ConsistOf(assertResponse()))
}

func assertResponse() []productCommand {
	return []productCommand{
		{
			Id:           "1",
			PriceInCents: 200000,
			Title:        "Notebook Gamer",
			Description:  "A great notebook",
			Discount: discount{
				Percentage:   5,
				ValueInCents: 10000,
			},
		},
		{
			Id:           "2",
			PriceInCents: 100000,
			Title:        "Phone Gamer",
			Description:  "A great phone",
			Discount: discount{
				Percentage:   5,
				ValueInCents: 5000,
			},
		},
		{
			Id:           "3",
			PriceInCents: 345000,
			Title:        "Monitor 21 pol",
			Description:  "A beautiful monitor",
			Discount: discount{
				Percentage:   5,
				ValueInCents: 17250,
			},
		},
		{
			Id:           "4",
			PriceInCents: 213000,
			Title:        "Notebook Dell",
			Description:  "The fastest dell notebook",
			Discount: discount{
				Percentage:   5,
				ValueInCents: 10650,
			},
		},
		{
			Id:           "5",
			PriceInCents: 879000,
			Title:        "MacBook Air",
			Description:  "The new MacBook Air",
			Discount: discount{
				Percentage:   5,
				ValueInCents: 43950,
			},
		},
		{
			Id:           "6",
			PriceInCents: 799000,
			Title:        "iPad Pro",
			Description:  "The new iPad",
			Discount: discount{
				Percentage:   5,
				ValueInCents: 39950,
			},
		},
	}
}

func fakeUser() domain.User {
	return domain.CreateUser("Nikola", "Tesla", time.Now().UTC().AddDate(-30, 0, 0))
}
