package domain

import (
	"github.com/ednailson/hash-challenge/discount-calculator/time_now"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	g := NewGomegaWithT(t)
	date := time.Date(2020, 11, 28, 19, 47, 30, 0, time.UTC)
	time_now.ReplaceFunctionTime(func() time.Time {
		return date
	})
	t.Run("creating a new user", func(t *testing.T) {
		sut := CreateUser("Albert", "Einstein", date.AddDate(-17, 0, 0))

		g.Expect(sut).Should(MatchAllFields(Fields{
			"Id":          BeEquivalentTo(""),
			"FirstName":   BeEquivalentTo("Albert"),
			"LastName":    BeEquivalentTo("Einstein"),
			"DateOfBirth": BeEquivalentTo(date.AddDate(-17, 0, 0)),
		}))
	})
	t.Run("it is birthday", func(t *testing.T) {
		sut := CreateUser("Albert", "Einstein", date.AddDate(-17, 0, 0))

		g.Expect(sut.IsBirthday()).Should(BeTrue())
	})
	t.Run("it is not birthday", func(t *testing.T) {
		sut := CreateUser("Albert", "Einstein", date.AddDate(0, -1, 0))

		g.Expect(sut.IsBirthday()).Should(BeFalse())
	})
}
