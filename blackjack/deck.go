package blackjack

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

func (d *Deck) Create() {
	for suit := 0; suit < NumSuits; suit++ {
		for value := MinCardValue; value <= MaxCardValue; value++ {
			d.Cards = append(d.Cards, Card{Value: value, Suit: Suit(suit)})
		}
	}
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}

func (d *Deck) Pop() Card {
	card := d.Cards[0]

	d.Cards = d.Cards[1:]
	return card
}

func (d Deck) PrintDeck() {
	for _, card := range d.Cards {
		fmt.Println(card.AsString())
	}
}
