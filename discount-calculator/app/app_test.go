package app

import (
	"crypto/tls"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/ednailson/hash-challenge/discount-calculator/database"
	"github.com/ednailson/hash-challenge/discount-calculator/domain"
	"github.com/ednailson/hash-challenge/discount-calculator/server/discount"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"strconv"
	"testing"
	"time"
)

func TestApp(t *testing.T) {
	g := NewGomegaWithT(t)
	sut, err := LoadApp(fakeConfig())
	g.Expect(err).ShouldNot(HaveOccurred())
	userColl := mockCollection(g, fakeConfig().Database.Config, fakeConfig().Database.UserCollection)
	productColl := mockCollection(g, fakeConfig().Database.Config, fakeConfig().Database.ProductCollection)
	userDocMeta, err := userColl.CreateDocument(nil, fakeUser())
	g.Expect(err).ShouldNot(HaveOccurred())
	productDocMeta, err := productColl.CreateDocument(nil, fakeProduct())
	g.Expect(err).ShouldNot(HaveOccurred())
	sut.Run()
	defer sut.Close()
	conn, client := grpcClient(g, fakeConfig().Port)
	defer conn.Close()

	assert, err := client.CalculateDiscount(context.Background(), &discount.Info{
		UserId:    userDocMeta.Key,
		ProductId: productDocMeta.Key,
	})

	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(assert).Should(PointTo(MatchFields(IgnoreExtras, Fields{
		"Percentage":   BeEquivalentTo(5),
		"ValueInCents": BeEquivalentTo(100),
	})))
}

func fakeConfig() Config {
	return Config{
		Port: 9000,
		Database: Database{
			Config: database.Config{
				Host:     "http://arangodb.service.com.br",
				Port:     8529,
				User:     "root",
				Password: "dummyPass",
				Database: "hash-db",
			},
			UserCollection:    "user-collection",
			ProductCollection: "product-collection",
		},
	}
}

func fakeUser() domain.User {
	return domain.CreateUser("Albert", "Einstein", time.Now().UTC().AddDate(-17, 0, 0))
}

func fakeProduct() domain.Product {
	return domain.CreateProduct(2000, "Notebook Gamer", "A great notebook")
}

func grpcClient(g *GomegaWithT, port int) (*grpc.ClientConn, discount.DiscountServiceClient) {
	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithInsecure(), grpc.WithBlock())
	g.Expect(err).ShouldNot(HaveOccurred())
	return conn, discount.NewDiscountServiceClient(conn)
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
