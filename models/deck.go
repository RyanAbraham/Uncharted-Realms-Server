package models

import (
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card `json:"cards"`
}

func (d *Deck) drawCard() Card {
	var c Card
	c, d.Cards = d.Cards[0], d.Cards[1:]
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
	// Code from https://www.calhoun.io/how-to-shuffle-arrays-and-slices-in-go/
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(d.Cards) > 0 {
		n := len(d.Cards)
		randIndex := r.Intn(n)
		d.Cards[n-1], d.Cards[randIndex] = d.Cards[randIndex], d.Cards[n-1]
		d.Cards = d.Cards[:n-1]
	}
}
