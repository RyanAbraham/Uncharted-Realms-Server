package models

import (
	"reflect"
)

// CardZone holds a list of cards in that zone
type CardZone struct {
	Cards []Card
}

// AddCard adds a given card to the CardZone
func (cz *CardZone) AddCard(c Card) {
	cz.Cards = append(cz.Cards, c)
}

// RemoveCard removes a given card from the CardZone
func (cz *CardZone) RemoveCard(c Card) {
	if i := cz.getIndexOfCard(c); i != -1 {
		// From https://github.com/golang/go/wiki/SliceTricks delete
		copy(cz.Cards[i:], cz.Cards[i+1:])
		cz.Cards[len(cz.Cards)-1] = Card{}
		cz.Cards = cz.Cards[:len(cz.Cards)-1]
	}
}

func (cz *CardZone) getIndexOfCard(c Card) int {
	for i, c2 := range cz.Cards {
		if reflect.DeepEqual(c, c2) {
			return i
		}
	}
	return -1
}
