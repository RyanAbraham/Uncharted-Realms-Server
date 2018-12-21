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
	seed      int64
}

// GameResult contains all information from a game that was simulated
type GameResult struct {
	Winner       int                  `json:"winner"`
	AnimationLog *models.AnimationLog `json:"animationLog"`
	Seed         int64                `json:"seed"`
}

var aniLog models.AnimationLog
var result GameResult
var f *os.File

// Play simulates a game with 2 decks.
// Returns 0 for player 1 win, 1 for player 2, and -1 for a draw.
func Play(d1, d2 models.Deck) GameResult {
	var state gameState
	startGame(&state, d1, d2)
	defer f.Close()
	log.Println("*** START GAME ***")
	for {
		playTurn(&state)
		// Check if a player won through damage
		if state.playerHPs[0] <= 0 && state.playerHPs[1] <= 0 {
			// Tie game
			log.Printf("*** GAME ENDED ***\nWINNER: DRAW\n\n")
			result.Winner = -1
			return result
		}
		for idx, hp := range state.playerHPs {
			if hp <= 0 {
				// The other player won
				log.Printf("*** GAME ENDED ***\nWINNER: PLAYER %d\n\n", (idx+1)%2+1)
				result.Winner = (idx + 1) % 2
				return result
			}
		}
	}
}

func playTurn(state *gameState) {
	p := state.turn   // Player turn
	ep := (p + 1) % 2 // Enemy turn
	aniLog.StartTurn(p)
	log.Printf("TURN: Player %d\n", p+1)

	drawCard(state)

	reduceClocksAndPlayCards(state)

	declareAttacks(state)

	// End turn
	state.turn = ep
}

func startGame(state *gameState, d1, d2 models.Deck) {
	// Start logging
	newPath := filepath.Join(".", "logs")
	os.MkdirAll(newPath, os.ModePerm)
	logName := time.Now().Format("2006-01-02_15-04-05") + ".log"
	path := filepath.Join(newPath, logName)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)

	// Set up game variables
	state.seed = time.Now().UTC().UnixNano()
	result.Seed = state.seed
	result.AnimationLog = &aniLog

	// Populate and shuffle decks
	state.decks = [...]models.Deck{d1, d2}
	for i, deck := range state.decks {
		aniLog.ShuffleDeck(i)
		deck.Shuffle(state.seed)
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

func drawCard(state *gameState) {
	p := state.turn
	if c, e := state.decks[p].DrawCard(); e == nil {
		aniLog.DrawCard(p)
		log.Printf("Player %d draws a card\n", p+1)
		state.hands[p].AddCard(c)
	}
}

func reduceClocksAndPlayCards(state *gameState) {
	p := state.turn
	logGamestate(state)
	removed := 0 // Number of cards removed from hand while iterating
	for idx := range state.hands[p].Cards {
		j := idx - removed
		c := state.hands[p].Cards[j]
		aniLog.CardClocksDown(p, j)
		if c.ClockDown() {
			aniLog.PlayCard(p, j)
			log.Printf("Player %d plays card %d from their hand\n", p+1, j+1)
			state.fields[p].AddCard(c)
			state.hands[p].RemoveCard(c)
			removed++
		}
	}
}

func declareAttacks(state *gameState) {
	p := state.turn
	ep := (p + 1) % 2
	logGamestate(state)
	for idx, c := range state.fields[p].Cards {
		aniLog.CardAttacks(p, idx)
		log.Printf("Player %d attacks with %+v\n", p+1, c)
		if len(state.fields[ep].Cards) > idx {
			// There is an enemy card blocking this attack
			ec := state.fields[ep].Cards[idx]
			aniLog.CardAttacked(ep, idx, c.Pow)
			log.Printf("Player %d's %+v is attacked for %d damage\n", ep+1, ec, c.Pow)
			killed := ec.Damage(c.Pow)
			if killed {
				aniLog.CardDies(p, idx)
				log.Printf("Player %d's %+v dies\n", ep+1, ec)
				state.fields[ep].RemoveCard(ec)
			}
		} else {
			aniLog.PlayerAttacked(ep, c.Pow)
			log.Printf("Player %d is attacked for %d damage\n", ep+1, c.Pow)
			state.playerHPs[ep] -= c.Pow
		}
	}
}
