package models

import (
	"fmt"
)

// AnimationLog is a list of animations that happened in the game
type AnimationLog struct {
	// TODO: Make Animations a linked list for faster push/pop
	Animations animationList `json:"animations"`
	RandomSeed int64         `json:"seed"`
}

type animationList []animation

// animationType is used for an enum to distinguish animations
type animationType int

// animationType enum
const (
	startTurn animationType = iota
	drawCard
	playCard
	cardClocksDown
	cardAttacks
	cardAttacked
	cardDies
	cardStatIncrease
	cardStatDecrease
	shuffleDeck
	playerAttacked
	playerGainsLife
)

// cardLocation is used for an enum to represent different card locations
type cardLocation int

// cardLocation enum
const (
	deckLoc cardLocation = iota
	handLoc
	fieldLoc
)

type animation struct {
	Type   animationType `json:"type"`
	Player int           `json:"p"`
	CIdx   int           `json:"cIdx"`
	Val    int           `json:"val"`
	Loc    cardLocation  `json:"loc"`
}

func (a *AnimationLog) addAnimation(ani animation) {
	a.Animations = append(a.Animations, ani)
}

// StartTurn adds an animation of that type to the animation log
func (a *AnimationLog) StartTurn(p int) {
	a.addAnimation(animation{
		Type:   startTurn,
		Player: p,
	})
}

// DrawCard adds an animation of that type to the animation log
func (a *AnimationLog) DrawCard(p int) {
	a.addAnimation(animation{
		Type:   drawCard,
		Player: p,
	})
}

// PlayCard adds an animation of that type to the animation log
func (a *AnimationLog) PlayCard(p, idx int) {
	a.addAnimation(animation{
		Type:   playCard,
		Player: p,
		CIdx:   idx,
	})
}

// CardClocksDown adds an animation of that type to the animation log
func (a *AnimationLog) CardClocksDown(p, idx int) {
	a.addAnimation(animation{
		Type:   cardClocksDown,
		Player: p,
		CIdx:   idx,
	})
}

// CardAttacks adds an animation of that type to the animation log
func (a *AnimationLog) CardAttacks(p, idx int) {
	a.addAnimation(animation{
		Type:   cardAttacks,
		Player: p,
		CIdx:   idx,
	})
}

// CardAttacked adds an animation of that type to the animation log
func (a *AnimationLog) CardAttacked(p, idx, val int) {
	a.addAnimation(animation{
		Type:   cardAttacked,
		Player: p,
		CIdx:   idx,
		Val:    val,
	})
}

// CardDies adds an animation of that type to the animation log
func (a *AnimationLog) CardDies(p, idx int) {
	a.addAnimation(animation{
		Type:   cardAttacked,
		Player: p,
		CIdx:   idx,
	})
}

// CardStatIncrease adds an animation of that type to the animation log
func (a *AnimationLog) CardStatIncrease(p, idx, val int, loc cardLocation) {
	a.addAnimation(animation{
		Type:   cardStatIncrease,
		Player: p,
		CIdx:   idx,
		Val:    val,
		Loc:    loc,
	})
}

// CardStatDecrease adds an animation of that type to the animation log
func (a *AnimationLog) CardStatDecrease(p, idx, val int, loc cardLocation) {
	a.addAnimation(animation{
		Type:   cardStatDecrease,
		Player: p,
		CIdx:   idx,
		Val:    val,
		Loc:    loc,
	})
}

// ShuffleDeck adds an animation of that type to the animation log
func (a *AnimationLog) ShuffleDeck(p int) {
	a.addAnimation(animation{
		Type:   shuffleDeck,
		Player: p,
	})
}

// PlayerAttacked adds an animation of that type to the animation log
func (a *AnimationLog) PlayerAttacked(p, val int) {
	a.addAnimation(animation{
		Type:   playerAttacked,
		Player: p,
		Val:    val,
	})
}

// PlayerGainsLife adds an animation of that type to the animation log
func (a *AnimationLog) PlayerGainsLife(p, val int) {
	a.addAnimation(animation{
		Type:   playerGainsLife,
		Player: p,
		Val:    val,
	})
}

func (aniT animationType) String() string {
	animationTypes := [...]string{
		"startTurn",
		"drawCard",
		"playCard",
		"cardClocksDown",
		"cardAttacks",
		"cardAttacked",
		"cardDies",
		"cardStatIncrease",
		"cardStatDecrease",
		"shuffleDeck",
		"playerAttacked",
		"playerGainsLife",
	}
	if int(aniT) < len(animationTypes) {
		return animationTypes[aniT]
	}
	return "unknown"
}

func (anis animationList) String() string {
	s := "[\n"
	for _, a := range anis {
		s += fmt.Sprintf("%+v\n", a)
	}
	return s + "]"
}
