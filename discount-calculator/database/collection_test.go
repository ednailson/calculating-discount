package database

import (
	"github.com/arangodb/go-driver"
	. "github.com/onsi/gomega"
	"testing"
)

const testCollection = "test-collection"

func TestCollection(t *testing.T) {
	g := NewGomegaWithT(t)
	db, err := NewDatabase(fakeDbConfig())
	g.Expect(err).ShouldNot(HaveOccurred())

	t.Run("validate the collection creation", func(t *testing.T) {
		sut, err := db.Collection(testCollection)

		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(sut).ShouldNot(BeNil())
		arangoClient := mockClient(g, fakeDbConfig())
		arangoClientDB, err := arangoClient.Database(nil, fakeDbConfig().Database)
		g.Expect(err).ShouldNot(HaveOccurred())
		exists, err := arangoClientDB.CollectionExists(nil, testCollection)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(exists).Should(BeTrue())
	})

	t.Run("reading a document by id", func(t *testing.T) {
		arangoColl := mockCollection(g, fakeDbConfig(), testCollection)
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

func mockCollection(g *GomegaWithT, config Config, collName string) driver.Collection {
	client := mockClient(g, config)
	db, err := client.Database(nil, config.Database)
	g.Expect(err).ToNot(HaveOccurred())
	coll, err := db.Collection(nil, collName)
	g.Expect(err).ToNot(HaveOccurred())
	err = coll.Truncate(nil)
	g.Expect(err).ToNot(HaveOccurred())
	return coll
}
