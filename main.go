package main

import (
	"encoding/json"
	"net/http"

	"github.com/ryanabraham/urserver/models"
)

type response struct {
	winner      string
	gameActions []gameAction
}

type gameAction struct {
	action      string
	cardIDs     []string
	cardChanges []string
}

type statModifiedAction struct {
	modifiedCardID int
	statModified   string
	amountModified int
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		sampleCard := models.Card{
			ID:   1,
			Name: "Dark Confidant",
			Clk:  1,
			Pow:  2,
			Hp:   1,
			Img:  "testurl.com",
			Eff:  "Charge",
		}
		sampleDeck := models.Deck{
			Cards: make([]models.Card, 0),
		}
		for x := 0; x < 30; x++ {
			sampleDeck.Cards = append(sampleDeck.Cards, sampleCard)
		}
		theResponse, _ := json.Marshal(sampleDeck)
		w.Write(theResponse)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/game", gameHandler)
	http.ListenAndServe(":8080", nil)
}
