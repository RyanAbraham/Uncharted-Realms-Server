package simulators

import (
	"github.com/ryanabraham/urserver/models"
)

type gameState struct {
	hands  [2]models.CardZone
	fields [2]models.CardZone
	decks  [2]models.Deck
}

var state gameState
var turn int

func startGame() {
}
