package blackjack_test

import (
	"blackjack/blackjack"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Context("CardTest", func() {
	Context("#ValueAsString", func() {
		var (
			card blackjack.Card
		)

		It("Returns the Value as a string", func() {
			card.Value = 9
			Expect(card.ValueAsString()).To(Equal("9"))
		})

		Describe("When the Value is 1", func() {
			It("Returns Ace", func() {
				card.Value = 1
				Expect(card.ValueAsString()).To(Equal("Ace"))
			})
		})

		Describe("When the Value is 11", func() {
			It("Returns Jack", func() {
				card.Value = 11
				Expect(card.ValueAsString()).To(Equal("Jack"))
			})
		})

		Describe("When the Value is 12", func() {
			It("Returns Queen", func() {
				card.Value = 12
				Expect(card.ValueAsString()).To(Equal("Queen"))

			})
		})

		Describe("When the Value is 13", func() {
			It("Returns King", func() {
				card.Value = 13
				Expect(card.ValueAsString()).To(Equal("King"))
			})
		})
	})

	Context("#SuitAsString", func() {
		var(
			suit blackjack.Suit
		)
		Describe("When the suit is a Spade", func() {
			It("Returns Spade", func() {
				suit = blackjack.Spade
				Expect(suit.AsString()).To(Equal("Spade"))
			})
		})

		Describe("When the suit is a Heart", func() {
			It("Returns Heart", func() {
				suit = blackjack.Heart
				Expect(suit.AsString()).To(Equal("Heart"))
			})
		})

		Describe("When the suit is a Club", func() {
			It("Returns Club", func() {
				suit = blackjack.Club
				Expect(suit.AsString()).To(Equal("Club"))
			})
		})

		Describe("When the suit is a Diamond", func() {
			It("Returns Diamond", func() {
				suit = blackjack.Diamond
				Expect(suit.AsString()).To(Equal("Diamond"))
			})
		})

	})

	Context("#AsString", func() {
		It("returns a string containing the value and the suit of the card", func() {
			card := blackjack.Card{
				Value: 13,
				Suit: blackjack.Diamond,
			}

			Expect(card.AsString()).To(Equal("King of Diamonds"))
		})
	})

	Context("Points", func() {
		var (
			card blackjack.Card
		)

		It("returns the value of the card", func() {
			card.Value = 7
			Expect(card.Points()).To(Equal(7))
		})

		Describe("when the card is an ace", func() {
			It("returns 11", func() {
				card.Value = 1
				Expect(card.Points()).To(Equal(11))
			})
		})

		Describe("when the card is a face card", func() {
			It("returns 10", func() {
				card.Value = 11
				Expect(card.Points()).To(Equal(10))
			})
		})
	})
})