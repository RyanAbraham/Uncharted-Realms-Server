package handlers

import (
	"net/http"

	"github.com/ryanabraham/urserver/helpers"
	"github.com/ryanabraham/urserver/simulators"
)

// DebugHandler serves a sample JSON response for a game
func (a *App) DebugHandler(w http.ResponseWriter, r *http.Request) {
	var one = helpers.FakeCard(1, 1, 1, nil)
	var two = helpers.FakeCard(2, 2, 2, nil)
	var oneDeck = helpers.DeckOf(one, 10)
	var twoDeck = helpers.DeckOf(two, 10)
	result := simulators.Play(twoDeck, oneDeck)
	RespondWithJSON(w, 200, result)
}
