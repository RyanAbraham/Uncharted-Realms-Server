package models

import (
	"errors"
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
func (d *Deck) DrawCard() (Card, error) {
	var c Card
	if len(d.cards) > 0 {
		c, d.cards = d.cards[0], d.cards[1:]
		return c, nil
	}
	return c, errors.New("Can't draw from empty deck")
}

// DrawCards draws multiple cards from the deck and returns them all
func (d *Deck) DrawCards(num int) []Card {
	var cards []Card
	for x := 0; x < num; x++ {
		if c, e := d.DrawCard(); e != nil {
			cards = append(cards, c)
		}
	}
	return cards
}

// Shuffle randomizes the order of cards in the deck
func (d *Deck) Shuffle(seed int64) {
	// TODO: Test if this can be set in game.go
	rand.Seed(seed)
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}
