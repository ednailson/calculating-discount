package database

import (
	. "github.com/ednailson/hash-challenge/discount-calculator/helper_tests"
	. "github.com/onsi/gomega"
	"testing"
)

func TestDatabase(t *testing.T) {
	g := NewGomegaWithT(t)

	sut, err := NewDatabase(FakeDbConfig())

	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(sut).ShouldNot(BeNil())
	arangoClient := MockClient(g)
	exists, err := arangoClient.DatabaseExists(nil, DBNameTest)
	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(exists).Should(BeTrue())
}

func FakeDbConfig() Config {
	return Config{
		Host:     DBHostTest,
		Port:     DBPortTest,
		User:     DBUserTest,
		Password: DBPassTest,
		Database: DBNameTest,
	}
}
