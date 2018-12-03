package simulators

import (
	"fmt"
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

// GameResult contains all information from a game that was simulated
type GameResult struct {
	Winner       int                  `json:"winner"`
	AnimationLog *models.AnimationLog `json:"animationLog"`
	Seed         int64                `json:"seed"`
}

var state gameState
var seed int64
var aniLog models.AnimationLog
var result GameResult
var f *os.File

// Play simulates a game with 2 decks.
// Returns 0 for player 1 win, 1 for player 2, and -1 for a draw.
func Play(d1, d2 models.Deck) GameResult {
	startGame(d1, d2)
	defer f.Close()
	log.Println("*** START GAME ***")
	for {
		playTurn()
		// Check if a player won through damage
		if state.playerHPs[0] <= 0 && state.playerHPs[1] <= 0 {
			// Tie game
			log.Printf("### Animation Log: %+v\n", aniLog)
			log.Printf("*** GAME ENDED ***\nWINNER: DRAW\n\n")
			result.Winner = -1
			return result
		}
		for idx, hp := range state.playerHPs {
			if hp <= 0 {
				// The other player won
				log.Printf("### Animation Log: %+v\n", aniLog)
				log.Printf("*** GAME ENDED ***\nWINNER: PLAYER %d\n\n", (idx+1)%2+1)
				result.Winner = (idx + 1) % 2
				return result
			}
		}
	}
}

func playTurn() {
	p := state.turn   // Player turn
	ep := (p + 1) % 2 // Enemy turn
	aniLog.StartTurn(p)

	// Draw a card
	if c, e := state.decks[p].DrawCard(); e == nil {
		aniLog.DrawCard(p)
		state.hands[p].AddCard(&c)
	}

	// Reduce clocks
	for idx, c := range state.hands[p].Cards {
		aniLog.CardClocksDown(p, idx)
		if c.ClockDown() {
			// TODO: May need to adjust idx
			aniLog.PlayCard(p, idx)
			state.fields[p].AddCard(c)
			state.hands[p].RemoveCard(c)
		}
	}

	// Declare attacks
	for idx, c := range state.fields[p].Cards {
		aniLog.CardAttacks(p, idx)
		if len(state.fields[ep].Cards) > idx {
			// There is an enemy card blocking this attack
			aniLog.CardAttacked(ep, idx, c.Pow)
			ec := state.fields[ep].Cards[idx]
			killed := ec.Damage(c.Pow)
			if killed {
				aniLog.CardDies(p, idx)
				state.fields[ep].RemoveCard(ec)
			}
		} else {
			aniLog.PlayerAttacked(ep, c.Pow)
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

	// Set up game variables
	seed = time.Now().UTC().UnixNano()
	fmt.Println("### seed", seed)
	result.Seed = seed
	fmt.Println("### result.Seed", result.Seed)
	result.AnimationLog = &aniLog

	// Populate and shuffle decks
	state.decks = [...]models.Deck{d1, d2}
	for i, deck := range state.decks {
		aniLog.ShuffleDeck(i)
		deck.Shuffle(seed)
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
