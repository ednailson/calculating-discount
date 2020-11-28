package database

import (
	"crypto/tls"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	. "github.com/onsi/gomega"
	"strconv"
	"testing"
)

const dbHostTest = "http://arangodb.service.com.br"
const dbNameTest = "hash-db"
const dbPassTest = "dummyPass"
const dbUserTest = "root"
const dbPortTest = 8529

func TestDatabase(t *testing.T) {
	g := NewGomegaWithT(t)

	sut, err := NewDatabase(fakeDbConfig())

	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(sut).ShouldNot(BeNil())
	arangoClient := mockClient(g, fakeDbConfig())
	exists, err := arangoClient.DatabaseExists(nil, dbNameTest)
	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(exists).Should(BeTrue())
}

func mockClient(g *GomegaWithT, config Config) driver.Client {
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

func fakeDbConfig() Config {
	return Config{
		Host:     dbHostTest,
		Port:     dbPortTest,
		User:     dbUserTest,
		Password: dbPassTest,
		Database: dbNameTest,
	}
}
