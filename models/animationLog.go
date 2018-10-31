package models

// AnimationLog is a list of animations that happened in the game
type AnimationLog struct {
	// TODO: Make Animations a linked list for faster push/pop
	Animations []animation `json:"animationLog"`
}

// AnimationType is used for an enum to distinguish animations
type AnimationType int

// AnimationType enum
const (
	DrawCard          AnimationType = 0
	CardClocksDown    AnimationType = 1
	PlayCard          AnimationType = 2
	CardAttacks       AnimationType = 3
	CardTakesDamage   AnimationType = 4
	PlayerTakesDamage AnimationType = 5
	PlayerGainsLife   AnimationType = 6
)

// CardLocation is used for an enum to represent different card locations
type CardLocation int

// CardLocation enum
const (
	Deck  CardLocation = 0
	Hand  CardLocation = 1
	Field CardLocation = 2
)

type animation struct {
	Type   AnimationType `json:"type"`
	Player int           `json:"player"`
	CIdx   int           `json:"cIdx"`
	CLoc   CardLocation  `json:"cLoc"`
	Value  int           `json:"value"`
}

func (a *AnimationLog) addAnimation(ani animation) {
	a.Animations = append(a.Animations, ani)
}
