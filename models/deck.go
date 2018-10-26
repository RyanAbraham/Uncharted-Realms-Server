package models

import (
	"math/rand"
)

// Deck holds a variable number of game cards
type Deck struct {
	cards []Card
}

// AddCards adds all cards passed in to the deck
func (d *Deck) AddCards(cards ...Card) {
	for _, c := range cards {
		d.cards = append(d.cards, c)
	}
}

// DrawCard draws a single card from the deck and returns it
func (d *Deck) DrawCard() Card {
	var c Card
	if len(d.cards) > 0 {
		c, d.cards = d.cards[0], d.cards[1:]
		return c
	}
	return nil // No card to draw
}

// DrawCards draws multiple cards from the deck and returns them all
func (d *Deck) DrawCards(num int) []Card {
	var cards []Card
	for x := 0; x < num; x++ {
		cards = append(cards, d.DrawCard())
	}
	return cards
}

// Shuffle randomizes the order of cards in the deck
func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}
