package models

import "fmt"

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

// Damage reduces the hp of a card and returns true if HP 0
func (c *Card) Damage(d int) bool {
	c.Hp -= d
	return c.Hp <= 0
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

func (c *Card) String() string {
	return fmt.Sprintf("Card{ID: %d, Name: %s, Clk: %d, Pow: %d, Hp: %d, Effs: %v}",
		c.ID, c.Name, c.Clk, c.Pow, c.Hp, c.Effs)
}
