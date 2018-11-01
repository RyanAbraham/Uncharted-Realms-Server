package models

// AnimationLog is a list of animations that happened in the game
type AnimationLog struct {
	// TODO: Make Animations a linked list for faster push/pop
	Animations []animation `json:"animations"`
	RandomSeed int         `json:"seed"`
}

// animationType is used for an enum to distinguish animations
type animationType int

// animationType enum
const (
	drawCard          animationType = 0
	playCard          animationType = 1
	cardClocksDown    animationType = 2
	cardAttacks       animationType = 3
	cardTakesDamage   animationType = 4
	cardStatIncrease  animationType = 5
	cardStatDecrease  animationType = 6
	playerTakesDamage animationType = 7
	playerGainsLife   animationType = 8
)

// cardLocation is used for an enum to represent different card locations
type cardLocation int

// cardLocation enum
const (
	deckLoc  cardLocation = 0
	handLoc  cardLocation = 1
	fieldLoc cardLocation = 2
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

// CardTakesDamage adds an animation of that type to the animation log
func (a *AnimationLog) CardTakesDamage(p, idx, val int) {
	a.addAnimation(animation{
		Type:   cardTakesDamage,
		Player: p,
		CIdx:   idx,
		Val:    val,
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

// PlayerTakesDamage adds an animation of that type to the animation log
func (a *AnimationLog) PlayerTakesDamage(p, val int) {
	a.addAnimation(animation{
		Type:   playerTakesDamage,
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
