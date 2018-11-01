package simulators

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/ryanabraham/urserver/models"
)

// TODO: Possibly abstract this to its own model
type gameState struct {
	decks     [2]models.Deck
	hands     [2]models.CardZone
	fields    [2]models.CardZone
	playerHPs [2]int
	turn      int
}

var state gameState
var f *os.File

// Play simulates a game with 2 decks.
// Returns 0 for player 1 win, 1 for player 2, and -1 for a draw.
func Play(d1, d2 models.Deck) int {
	startGame(d1, d2)
	defer f.Close()
	for {
		log.Printf("***PLAYER %d TURN START***", state.turn+1)
		playTurn()
		log.Printf("state.hands %+v", state.hands)
		log.Printf("state.fields %+v", state.fields)
		log.Println("state.playerHPs", state.playerHPs)
		// Check if a player won through damage
		if state.playerHPs[0] <= 0 && state.playerHPs[1] <= 0 {
			return -1
		}
		for idx, hp := range state.playerHPs {
			if hp <= 0 {
				// The other player won
				return (idx + 1) % 2
			}
		}
	}
}

func playTurn() {
	p := state.turn   // Player turn
	ep := (p + 1) % 2 // Enemy turn

	// Draw a card
	if c, e := state.decks[p].DrawCard(); e == nil {
		state.hands[p].AddCard(&c)
	}

	// Reduce clocks
	for _, c := range state.hands[p].Cards {
		if c.ClockDown() {
			state.fields[p].AddCard(c)
			state.hands[p].RemoveCard(c)
		}
	}

	// Declare attacks
	for idx, c := range state.fields[p].Cards {
		if len(state.fields[ep].Cards) > idx {
			// There is an enemy card blocking this attack
			ec := state.fields[ep].Cards[idx]
			killed := ec.Damage(c.Pow)
			if killed {
				state.fields[ep].RemoveCard(ec)
			}
		} else {
			state.playerHPs[ep] -= c.Pow
		}
	}

	// End turn
	state.turn = ep
}

func startGame(d1, d2 models.Deck) {
	// Start logging
	newPath := filepath.Join("..", "logs")
	os.MkdirAll(newPath, os.ModePerm)
	logName := time.Now().Format("2006-01-02_15-04-05") + ".log"
	path := filepath.Join(newPath, logName)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)

	// Populate and shuffle decks
	state.decks = [...]models.Deck{d1, d2}
	for _, deck := range state.decks {
		deck.Shuffle()
	}
	// Initialize hands and fields
	state.hands = [2]models.CardZone{
		models.CardZone{},
		models.CardZone{},
	}
	state.fields = [2]models.CardZone{
		models.CardZone{},
		models.CardZone{},
	}

	// Starting player HPs
	for idx := range state.playerHPs {
		state.playerHPs[idx] = 30
	}
	// Set starting player
	// TODO: Randomize this
	state.turn = 0
}
