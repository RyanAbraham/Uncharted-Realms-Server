package models

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
)

// Deck holds a variable number of game cards
type Deck struct {
	Id    int    `json:"id"`
	Cards []Card `json:"cards"`
}

// AddCards adds all cards passed in to the deck
func (d *Deck) AddCards(cards ...Card) {
	for _, c := range cards {
		d.Cards = append(d.Cards, c)
	}
}

// DrawCard draws a single card from the deck and returns it
func (d *Deck) DrawCard() (Card, error) {
	var c Card
	if len(d.Cards) > 0 {
		c, d.Cards = d.Cards[0], d.Cards[1:]
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
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck) getDeck(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (d *Deck) updateDeck(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (d *Deck) deleteDeck(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (d *Deck) createDeck(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getDecks(db *sql.DB, start, count int) ([]Deck, error) {
	return nil, errors.New("Not implemented")
}

func (d Deck) String() string {
	str := fmt.Sprintf("Deck{Cards: ")
	for _, c := range d.Cards {
		str += fmt.Sprintf("\n  %+v", c)
	}
	return str + "\n}"
}
