package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	httplogger "github.com/jesseokeya/go-httplogger"
)

// App holds references to the router and database
type App struct {
	Router           *mux.Router
	DB               *sql.DB
	MatchmakingQueue chan string
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/debug", a.DebugHandler).Methods("GET")
}

// TryMatchmaking tries to match any users waiting for a game
func (a *App) TryMatchmaking() {
	userID := <-a.MatchmakingQueue
	userID2 := <-a.MatchmakingQueue
	if userID != "" && userID2 != "" {
		// Start a game with the two players

	}
}

// Initialize sets up the app by intializing the router and its routes
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

// Run starts the server at the given address
func (a *App) Run(addr string) {
	fmt.Println("Listening at", addr)
	http.ListenAndServe(addr, httplogger.Golog(a.Router))
}

// RespondWithError sends an error response as JSON
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJSON(w, code, map[string]string{"error": msg})
}

// RespondWithJSON sends a response to the client with the given parameters
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
