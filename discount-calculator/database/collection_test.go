package database

import (
	"github.com/ednailson/hash-challenge/discount-calculator/domain"
	. "github.com/ednailson/hash-challenge/discount-calculator/helper_tests"
	"github.com/ednailson/hash-challenge/discount-calculator/time_now"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"testing"
)

const testCollection = "test-collection"

func TestCollection(t *testing.T) {
	g := NewGomegaWithT(t)
	db, err := NewDatabase(FakeDbConfig())
	g.Expect(err).ShouldNot(HaveOccurred())

	t.Run("validate the collection creation", func(t *testing.T) {
		sut, err := db.Collection(testCollection)

		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(sut).ShouldNot(BeNil())
		arangoClient := MockClient(g)
		arangoClientDB, err := arangoClient.Database(nil, FakeDbConfig().Database)
		g.Expect(err).ShouldNot(HaveOccurred())
		exists, err := arangoClientDB.CollectionExists(nil, testCollection)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(exists).Should(BeTrue())
	})

	t.Run("reading a document by id", func(t *testing.T) {
		arangoColl := MockAndTruncateCollection(g, testCollection)
		date := time_now.Now()
		user := domain.CreateUser("Albert", "Einstein", date)
		docCreated, err := arangoColl.CreateDocument(nil, user)
		g.Expect(err).ShouldNot(HaveOccurred())
		coll, err := db.Collection(testCollection)
		g.Expect(err).ShouldNot(HaveOccurred())

		var assert domain.User
		err = coll.ReadById(docCreated.Key, &assert)

		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(assert).Should(MatchAllFields(Fields{
			"Id":          BeEquivalentTo(docCreated.Key),
			"FirstName":   BeEquivalentTo("Albert"),
			"LastName":    BeEquivalentTo("Einstein"),
			"DateOfBirth": BeEquivalentTo(date),
		}))
	})
}
