package simulators

import "log"

func logGamestate(state gameState) {
	log.Printf(`GAMESTATE:
		P1 Hand: %+v
		P1 Field: %+v
		P2 Field: %+v
		P2 Hand: %+v\n`,
		state.hands[0],
		state.fields[0],
		state.fields[1],
		state.hands[1])
}
