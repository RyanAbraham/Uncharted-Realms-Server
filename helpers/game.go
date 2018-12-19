package helpers

import "github.com/ryanabraham/urserver/models"

// FakeCard returns a card with a dummy name and ID with the passed in stats
func FakeCard(clk, pow, hp int, effs []string) *models.Card {
	return &models.Card{
		ID:   0,
		Name: "Test",
		Clk:  clk,
		Pow:  pow,
		Hp:   hp,
		Img:  "",
		Effs: effs,
	}
}

// DeckOf returns a Deck of given size with only copies of the given card
func DeckOf(c *models.Card, x int) models.Deck {
	fakeDeck := models.Deck{}
	for i := 0; i < x; i++ {
		newCard := c // Copy the object
		fakeDeck.AddCards(newCard)
	}
	return fakeDeck
}
