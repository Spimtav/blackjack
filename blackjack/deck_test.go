package blackjack_test

import (
	"blackjack/blackjack"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Context("#DeckTest", func () {
	Describe("When a deck is created", func() {
		It("Creates 52 card objects", func() {
			var deck blackjack.Deck
			deck.Create()
			Expect(len(deck.Cards)).To(Equal(52))
		})

		It("contains no duplicated cards", func() {
			cards := map[string]map[int]int{}

			var deck blackjack.Deck
			deck.Create()

			for _, card := range deck.Cards {
				_, ok := cards[card.Suit.AsString()][card.Value]
				if !ok {
					cards[card.Suit.AsString()] = map[int]int{}
				}

				cards[card.Suit.AsString()][card.Value]+= 1
				Expect(cards[card.Suit.AsString()][card.Value]).To(Equal(1))
			}
		})
	})

	Describe("When a deck is shuffled", func() {
		It("randomizes the order of the cards in the deck", func() {
			var deck blackjack.Deck
			deck.Create()
			beforeOrder := deck.AsString()
			deck.Shuffle()

			Expect(beforeOrder).NotTo(Equal(deck.AsString()))
		})
	})

	Describe("When a deck is Popped", func() {
		It("returns the first card in the deck", func() {
			var deck blackjack.Deck
			deck.Create()

			Expect(deck.Pop().AsString()).To(Equal("Ace of Spades"))
		})

		It("removes that card from the deck", func() {
			var deck blackjack.Deck
			deck.Create()
			poppedCard := deck.Pop()

			Expect(len(deck.Cards)).To(Equal(51))

			for _, card := range deck.Cards {
				Expect(card.AsString()).NotTo(Equal(poppedCard.AsString()))
			}
		})
	})
})
