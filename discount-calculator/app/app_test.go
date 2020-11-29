package app

import (
	"fmt"
	"github.com/ednailson/hash-challenge/discount-calculator/database"
	"github.com/ednailson/hash-challenge/discount-calculator/domain"
	. "github.com/ednailson/hash-challenge/discount-calculator/helper_tests"
	"github.com/ednailson/hash-challenge/discount-calculator/server/discount"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestApp(t *testing.T) {
	g := NewGomegaWithT(t)
	sut, err := LoadApp(fakeConfig())
	g.Expect(err).ShouldNot(HaveOccurred())
	userColl := MockCollection(g, fakeConfig().Database.UserCollection)
	productColl := MockCollection(g, fakeConfig().Database.ProductCollection)
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
				Host:     DBHostTest,
				Port:     DBPortTest,
				User:     DBUserTest,
				Password: DBPassTest,
				Database: DBNameTest,
			},
			UserCollection:    UserCollection,
			ProductCollection: ProductCollection,
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
