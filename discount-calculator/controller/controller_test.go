package controller

import (
	"crypto/tls"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/ednailson/hash-challenge/discount-calculator/database"
	"github.com/ednailson/hash-challenge/discount-calculator/domain"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"strconv"
	"testing"
	"time"
)

const dbHostTest = "http://arangodb.service.com.br"
const dbNameTest = "hash-db"
const dbPassTest = "dummyPass"
const dbUserTest = "root"
const dbPortTest = 8529
const productCollection = "product-collection"
const userCollection = "user-collection"

func TestController(t *testing.T) {
	g := NewGomegaWithT(t)
	db, err := database.NewDatabase(fakeDbConfig())
	g.Expect(err).ShouldNot(HaveOccurred())
	userColl, err := db.Collection(userCollection)
	g.Expect(err).ShouldNot(HaveOccurred())
	productColl, err := db.Collection(productCollection)
	g.Expect(err).ShouldNot(HaveOccurred())
	sut := NewController(userColl, productColl)
	g.Expect(err).ShouldNot(HaveOccurred())
	t.Run("calculating a discount", func(t *testing.T) {
		userColl := mockCollection(g, fakeDbConfig(), userCollection)
		productColl := mockCollection(g, fakeDbConfig(), productCollection)
		userDocMeta, err := userColl.CreateDocument(nil, fakeUser())
		g.Expect(err).ShouldNot(HaveOccurred())
		productDocMeta, err := productColl.CreateDocument(nil, fakeProduct())
		g.Expect(err).ShouldNot(HaveOccurred())

		productCommand, err := sut.CalculateDiscount(userDocMeta.Key, productDocMeta.Key)

		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(*productCommand).Should(MatchAllFields(Fields{
			"Id":           BeEquivalentTo(productDocMeta.Key),
			"Title":        BeEquivalentTo("Notebook Gamer"),
			"Description":  BeEquivalentTo("A great notebook"),
			"PriceInCents": BeEquivalentTo(2000),
			"Discount": MatchAllFields(Fields{
				"Percentage":   BeEquivalentTo(5),
				"ValueInCents": BeEquivalentTo(100),
			},
			),
		}))
	})
}

func mockCollection(g *GomegaWithT, config database.Config, collName string) driver.Collection {
	client := mockClient(g, config)
	db, err := client.Database(nil, config.Database)
	g.Expect(err).ToNot(HaveOccurred())
	coll, err := db.Collection(nil, collName)
	g.Expect(err).ToNot(HaveOccurred())
	err = coll.Truncate(nil)
	g.Expect(err).ToNot(HaveOccurred())
	return coll
}

func mockClient(g *GomegaWithT, config database.Config) driver.Client {
	dbConn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{config.Host + ":" + strconv.Itoa(config.Port)},
		TLSConfig: &tls.Config{},
	})
	g.Expect(err).ToNot(HaveOccurred())
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     dbConn,
		Authentication: driver.BasicAuthentication(config.User, config.Password)})
	g.Expect(err).ToNot(HaveOccurred())
	return client
}

func fakeDbConfig() database.Config {
	return database.Config{
		Host:     dbHostTest,
		Port:     dbPortTest,
		User:     dbUserTest,
		Password: dbPassTest,
		Database: dbNameTest,
	}
}

func fakeUser() domain.User {
	return domain.CreateUser("Albert", "Einstein", time.Now().UTC().AddDate(-17, 0, 0))
}

func fakeProduct() domain.Product {
	return domain.CreateProduct(2000, "Notebook Gamer", "A great notebook")
}
