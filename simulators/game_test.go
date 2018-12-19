package simulators

import (
	"reflect"
	"testing"

	"github.com/ryanabraham/urserver/helpers"
	"github.com/ryanabraham/urserver/models"
)

func TestGuaranteedWins(t *testing.T) {
	one := helpers.FakeCard(1, 1, 1, nil)
	oneDeck := helpers.DeckOf(one, 10)
	emptyDeck := models.Deck{}

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
	one := helpers.FakeCard(1, 1, 1, nil)
	oneDeck := helpers.DeckOf(one, 10)
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
	one := helpers.FakeCard(1, 1, 1, nil)
	two := helpers.FakeCard(2, 2, 2, nil)
	oneDeck := helpers.DeckOf(one, 10)
	twoDeck := helpers.DeckOf(two, 10)
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
	one := helpers.FakeCard(1, 1, 1, nil)
	oneDeck := helpers.DeckOf(one, 10)
	result := Play(oneDeck, oneDeck)
	if result.Winner != 1 {
		t.Errorf("Game result incorrect, got: %d, expected %d.", result, 1)
	}
}

func TestDrawCard(t *testing.T) {
	one := helpers.FakeCard(1, 1, 1, nil)
	state := gameState{
		decks: [2]models.Deck{
			models.Deck{Cards: []*models.Card{one}},
			models.Deck{},
		},
		turn: 0,
	}
	expectedState := state
	expectedState.decks[0].Cards = []*models.Card{}
	expectedState.hands[0].Cards = []*models.Card{one}

	drawCard(&state)

	if !reflect.DeepEqual(state, expectedState) {
		t.Errorf("Game state incorrect, got: %+v, expected %+v.", state, expectedState)
	}
}

func TestReduceClocksAndPlayCards(t *testing.T) {
	one := helpers.FakeCard(1, 1, 1, nil)
	two := helpers.FakeCard(2, 2, 2, nil)
	state := gameState{
		hands: [2]models.CardZone{
			models.CardZone{Cards: []*models.Card{one, two, one}},
			models.CardZone{},
		},
		turn: 0,
	}
	expectedState := state
	expectedState.hands[0].Cards = []*models.Card{helpers.FakeCard(1, 2, 2, nil)}
	expectedState.fields[0].Cards = []*models.Card{
		helpers.FakeCard(-1, 1, 1, nil),
		helpers.FakeCard(-1, 1, 1, nil),
	}

	reduceClocksAndPlayCards(&state)

	if !reflect.DeepEqual(state, expectedState) {
		t.Errorf("Game state incorrect, got: %+v, expected %+v.", state, expectedState)
	}
}
