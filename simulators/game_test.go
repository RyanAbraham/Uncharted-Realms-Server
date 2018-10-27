package simulators

import (
	"testing"

	"github.com/ryanabraham/urserver/helpers"
	"github.com/ryanabraham/urserver/models"
)

var one = helpers.FakeCard(1, 1, 1, nil)
var oneDeck = helpers.DeckOf(one, 10)
var emptyDeck = models.Deck{}

func TestGuaranteedWins(t *testing.T) {
	result := Play(oneDeck, emptyDeck)
	if result != 0 {
		t.Errorf("Game result incorrect, got: %d, expected %d.", result, 0)
	}
	result = Play(emptyDeck, oneDeck)
	if result != 1 {
		t.Errorf("Game result incorrect, got: %d, expected %d.", result, 1)
	}
}
func TestMoreCardsWins(t *testing.T) {
	oneMinus := helpers.DeckOf(one, 9)
	result := Play(oneDeck, oneMinus)
	if result != 0 {
		t.Errorf("Game result incorrect, got: %d, expected %d.", result, 0)
	}
	result = Play(oneMinus, oneDeck)
	if result != 1 {
		t.Errorf("Game result incorrect, got: %d, expected %d.", result, 1)
	}
}
