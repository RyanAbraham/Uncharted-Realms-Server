package handlers

import (
	"net/http"

	"github.com/ryanabraham/urserver/helpers"
	"github.com/ryanabraham/urserver/models"
	"github.com/ryanabraham/urserver/simulators"
)

type debugResponse struct {
	Winner       int                  `json:"winner"`
	AnimationLog *models.AnimationLog `json:"animationLog"`
	Seed         int64                `json:"seed"`
	Decks        []models.Deck        `json:"decks"`
}

// DebugHandler serves a sample JSON response for a game
func (a *App) DebugHandler(w http.ResponseWriter, r *http.Request) {
	var oneDeck = helpers.DeckOf(helpers.FakeCard(1, 1, 1, nil), 10)
	var twoDeck = helpers.DeckOf(helpers.FakeCard(2, 2, 2, nil), 10)
	result := simulators.Play(twoDeck, oneDeck)
	response := debugResponse{
		Winner:       result.Winner,
		AnimationLog: result.AnimationLog,
		Seed:         result.Seed,
		Decks:        []models.Deck{oneDeck, twoDeck},
	}
	RespondWithJSON(w, 200, response)
}
