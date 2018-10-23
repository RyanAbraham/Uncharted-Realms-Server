package models

import (
	"math/rand"
)

type Deck struct {
	cards []Card
}

func (d *Deck) drawCard() Card {
	var c Card
	c, d.cards = d.cards[0], d.cards[1:]
	return c
}

func (d *Deck) drawCards(num int) []Card {
	var cards []Card
	for x := 0; x < num; x++ {
		cards = append(cards, d.drawCard())
	}
	return cards
}

func (d *Deck) shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}
