package controller

import (
	"github.com/ednailson/hash-challenge/discount-calculator/database"
	"github.com/ednailson/hash-challenge/discount-calculator/domain"
	. "github.com/ednailson/hash-challenge/discount-calculator/helper_tests"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"testing"
	"time"
)

func TestController(t *testing.T) {
	g := NewGomegaWithT(t)
	db, err := database.NewDatabase(fakeDbConfig())
	g.Expect(err).ShouldNot(HaveOccurred())
	userColl, err := db.Collection(UserCollection)
	g.Expect(err).ShouldNot(HaveOccurred())
	productColl, err := db.Collection(ProductCollection)
	g.Expect(err).ShouldNot(HaveOccurred())
	sut := NewController(userColl, productColl)
	g.Expect(err).ShouldNot(HaveOccurred())
	t.Run("calculating a discount", func(t *testing.T) {
		userCollection := MockCollection(g, UserCollection)
		productCollection := MockCollection(g, ProductCollection)
		userDocMeta, err := userCollection.CreateDocument(nil, fakeUser())
		g.Expect(err).ShouldNot(HaveOccurred())
		productDocMeta, err := productCollection.CreateDocument(nil, fakeProduct())
		g.Expect(err).ShouldNot(HaveOccurred())

		productCommand, err := sut.CalculateDiscount(userDocMeta.Key, productDocMeta.Key)

		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(*productCommand).Should(MatchAllFields(Fields{
			"Percentage":   BeEquivalentTo(5),
			"ValueInCents": BeEquivalentTo(100),
		}))
	})
}

func fakeDbConfig() database.Config {
	return database.Config{
		Host:     DBHostTest,
		Port:     DBPortTest,
		User:     DBUserTest,
		Password: DBPassTest,
		Database: DBNameTest,
	}
}

func fakeUser() domain.User {
	return domain.CreateUser("Albert", "Einstein", time.Now().UTC().AddDate(-17, 0, 0))
}

func fakeProduct() domain.Product {
	return domain.CreateProduct(2000, "Notebook Gamer", "A great notebook")
}
