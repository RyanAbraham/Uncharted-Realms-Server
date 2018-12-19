package simulators

import (
	"fmt"
	"log"
)

func logGamestate(state *gameState) {
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

func (s gameState) String() string {
	str := "State: {\n"
	str += "  decks: {\n"
	for _, d := range s.decks {
		str += fmt.Sprintf("    %+v\n", d)
	}
	str += "  }\n"
	str += "  hands: {\n"
	for _, h := range s.hands {
		str += fmt.Sprintf("    %+v\n", h)
	}
	str += "  }\n"
	str += "  fields: {\n"
	for _, f := range s.fields {
		str += fmt.Sprintf("    %+v\n", f)
	}
	str += "  }\n"
	str += "}"
	return str
}
