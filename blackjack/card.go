package blackjack

import (
	"fmt"
	"math"
	"strconv"
)

type Suit int

const (
	Spade   Suit = iota
	Heart        = iota
	Club         = iota
	Diamond      = iota
)

func (s Suit) AsString() string {
	var suitName string
	switch s {
	case Spade:
		suitName = "Spade"
	case Heart:
		suitName = "Heart"
	case Club:
		suitName = "Club"
	case Diamond:
		suitName = "Diamond"
	}

	return suitName
}


const (
	NumSuits = 4

	MinCardValue = 1
	MaxCardValue = 13
)

type Card struct {
	Value int
	Suit  Suit
}

func (c Card) ValueAsString() string {
	switch c.Value {
	case 1:
		return "Ace"
	case 11:
		return "Jack"
	case 12:
		return "Queen"
	case 13:
		return "King"
	default:
		return strconv.Itoa(c.Value)
	}
}

func (c Card) AsString() string {
	value := c.ValueAsString()
	suit := fmt.Sprintf("%ss", c.Suit.AsString())

	return fmt.Sprintf("%s of %s", value, suit)
}

func (c Card) Points() int {
	if c.Value == 1 {
		return 11
	}

	points := int(math.Min(float64(c.Value), 10.0))

	return points
}
