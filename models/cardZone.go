package models

import (
	"reflect"
)

// CardZone holds a list of cards in that zone
type CardZone struct {
	cards []Card
}

// AddCard adds a given card to the CardZone
func (cz *CardZone) AddCard(c Card) {
	cz.cards = append(cz.cards, c)
}

// RemoveCard removes a given card from the CardZone
func (cz *CardZone) RemoveCard(c Card) {
	if i := cz.getIndexOfCard(c); i != -1 {
		// From https://github.com/golang/go/wiki/SliceTricks delete
		copy(cz.cards[i:], cz.cards[i+1:])
		cz.cards[len(cz.cards)-1] = Card{}
		cz.cards = cz.cards[:len(cz.cards)-1]
	}
}

func (cz *CardZone) getIndexOfCard(c Card) int {
	for i, c2 := range cz.cards {
		if reflect.DeepEqual(c, c2) {
			return i
		}
	}
	return -1
}
