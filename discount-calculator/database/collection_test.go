package database

import (
	. "github.com/ednailson/hash-challenge/discount-calculator/helper_tests"
	. "github.com/onsi/gomega"
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
		arangoColl := MockCollection(g, testCollection)
		document := map[string]string{"name": "Albert", "nickname": "Einstein"}
		docCreated, err := arangoColl.CreateDocument(nil, document)
		g.Expect(err).ShouldNot(HaveOccurred())
		coll, err := db.Collection(testCollection)
		g.Expect(err).ShouldNot(HaveOccurred())

		sut, err := coll.ReadById(docCreated.Key)

		g.Expect(err).ShouldNot(HaveOccurred())
		assert, ok := sut.(map[string]interface{})
		g.Expect(ok).Should(BeTrue())
		g.Expect(assert["name"]).Should(BeEquivalentTo(document["name"]))
		g.Expect(assert["nickname"]).Should(BeEquivalentTo(document["nickname"]))
	})
}
