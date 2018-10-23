package models

// Card represents a single unique game card
type Card struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Clk  int    `json:"clk"`
	Pow  int    `json:"pow"`
	Hp   int    `json:"hp"`
	Img  string `json:"img"`
	Eff  string `json:"eff"`
}

// Damage reduces the hp of a card
func (c *Card) Damage(d int) {
	c.Hp -= d
}

// ChangePow modifies the power of a card
func (c *Card) ChangePow(x int) {
	c.Pow += x
}

// ChangeClk changes the clock of a card and returns
// whether the card's clock has hit 0 or not
func (c *Card) ChangeClk(x int) bool {
	c.Clk += x
	return c.Clk <= 0
}

// ClockDown decrements the clock of a card and returns
// whether the card's clock has hit 0 or not
func (c *Card) ClockDown() bool {
	return c.ChangeClk(-1)
}
