package handlers

import (
	"net/http"

	"github.com/ryanabraham/urserver/models"
)

func (a *App) MatchmakingHandler(w http.ResponseWriter, r *http.Request) {
	// Where does user come from? Get id from body? or user obj from somewhere?
	var u models.User
	// decoder := json.NewDecoder(r.Body)
	a.MatchmakingQueue <- u.ID
	RespondWithJSON(w, 200, nil) // TODO: Payload
}
