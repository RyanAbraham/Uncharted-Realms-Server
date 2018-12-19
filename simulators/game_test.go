package simulators

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ryanabraham/urserver/helpers"
	"github.com/ryanabraham/urserver/models"
)

var one = helpers.FakeCard(1, 1, 1, nil)
var two = helpers.FakeCard(2, 2, 2, nil)
var oneDeck = helpers.DeckOf(one, 10)
var twoDeck = helpers.DeckOf(two, 10)
var emptyDeck = models.Deck{}

func TestGuaranteedWins(t *testing.T) {
	result := Play(oneDeck, emptyDeck)
	if result.Winner != 0 {
		t.Errorf("Game result incorrect, got: %d, expected %d.", result, 0)
	}
	result = Play(emptyDeck, oneDeck)
	if result.Winner != 1 {
		t.Errorf("Game result incorrect, got: %d, expected %d.", result, 1)
	}
}

func TestMoreCardsWins(t *testing.T) {
	oneMinusDeck := helpers.DeckOf(one, 9)
	result := Play(oneDeck, oneMinusDeck)
	if result.Winner != 0 {
		t.Errorf("Game result incorrect, got: %d, expected %d.", result, 0)
	}
	result = Play(oneMinusDeck, oneDeck)
	if result.Winner != 1 {
		t.Errorf("Game result incorrect, got: %d, expected %d.", result, 1)
	}
}

func TestStrongerCardsWins(t *testing.T) {
	result := Play(twoDeck, oneDeck)
	if result.Winner != 0 {
		t.Errorf("Game result incorrect, got: %d, expected %d.", result, 0)
	}
	result = Play(oneDeck, twoDeck)
	if result.Winner != 1 {
		t.Errorf("Game result incorrect, got: %d, expected %d.", result, 1)
	}
}

func TestEqualDeckWinsOnDraw(t *testing.T) {
	result := Play(oneDeck, oneDeck)
	if result.Winner != 1 {
		t.Errorf("Game result incorrect, got: %d, expected %d.", result, 1)
	}
}

func TestDrawCard(t *testing.T) {
	state := gameState{
		decks: [2]models.Deck{
			models.Deck{Cards: []models.Card{one}},
			models.Deck{},
		},
	}
	expectedState := state
	expectedState.decks[0].Cards = []models.Card{}
	expectedState.hands[0].Cards = []*models.Card{&one}
	drawCard(&state, 0)
	fmt.Println(state.hands[0])
	if !reflect.DeepEqual(state, expectedState) {
		t.Errorf("Game state incorrect, got: %+v, expected %+v.", state, expectedState)
	}
}
