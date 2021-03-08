package main

import (
	"blackjack/blackjack"
	"fmt"
)

func main() {
	fmt.Println("lets make a deck")

	var deck blackjack.Deck
	deck.Create()
	deck.PrintDeck()

	fmt.Println("\n\n\nOk NOW its shuffled-aru")
	deck.Shuffle()
	deck.PrintDeck()
}




/*
Classes:
- card
  - types:
    - Suit - enum for suit
  - struct:
    - int number between 1-11 (aces high)
    - a Suit
- deck
  - struct:
    - []Card slice of all 52 unique cards
  - funcs:
    - create()
    - shuffle()
    - pop() Card
- hand
  - struct:
    - []Card
  - funcs:
    - addCard(Card)
    - score() int
    - isBust() bool
- game
  - struct:
    - Deck
    - Hand (player)
    - Hand (dealer)
  - funcs:
    - initialize
    - askPlayerHitOrStay
    - isGameOver
    - whoWon() string

After:
- Player
  - struct:
    - string name
- casino/runner
  -



*/