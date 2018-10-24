package simulators

import (
	"github.com/ryanabraham/urserver/models"
)

// TODO: Possibly abstract this to its own model
type gameState struct {
	hands     [2]models.CardZone
	fields    [2]models.CardZone
	decks     [2]models.Deck
	playerHPs [2]int
	turn      int
}

var state gameState

// PlayGame simulates a game with 2 decks.
// Returns 1 for player 1 win, 2 for player 2, and 0 for a draw.
func PlayGame(d1, d2 models.Deck) int {
	startGame(d1, d2)
	for {
		playTurn()
		// Check if a player won through damage
		for idx, hp := range state.playerHPs {
			if hp <= 0 {
				return idx + 1
			}
		}
	}
	return 0
}

func playTurn() {
	p := state.turn   // Player turn
	ep := (p + 1) % 2 // Enemy turn

	// Draw a card
	state.hands[p].AddCard(state.decks[p].DrawCard())

	// Reduce clocks
	for _, c := range state.hands[p].Cards {
		if c.ClockDown() {
			state.fields[p].AddCard(c)
			state.hands[p].RemoveCard(c)
		}
	}

	// Declare attacks
	for idx, c := range state.fields[p].Cards {
		if len(state.fields[ep] >= idx {
			// There is an enemy card blocking this attack
			state.fields[ep].Cards[idx].Damage(c.Pow)
		} else {
			state.playerHPs[ep] -= c.Pow
		}
	}
	
	// End turn
	state.turn = ep
}

func startGame(d1, d2 models.Deck) {
	state.decks = [...]models.Deck{d1, d2}
	for _, deck := range state.decks {
		deck.Shuffle()
	}
	state.turn = 0
}
