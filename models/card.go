package models

// Card represents a single unique game card
type Card struct {
	ID   int      `json:"id"`
	Name string   `json:"name"`
	Clk  int      `json:"clk"`
	Pow  int      `json:"pow"`
	Hp   int      `json:"hp"`
	Img  string   `json:"img"`
	Effs []string `json:"effs"`
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

// AddEffs adds effects to a card
func (c *Card) AddEffs(effs ...string) {
	for _, eff := range effs {
		c.Effs = append(c.Effs, eff)
	}
}

// ClockDown decrements the clock of a card and returns
// whether the card's clock has hit 0 or not
func (c *Card) ClockDown() bool {
	return c.ChangeClk(-1)
}
