package deck

import "github.com/alanxoc3/concards-go/card"

// Deck contains a set of cards
type Deck struct {
	Cards []card.Card
}

// Open opens filename and loads cards into new deck
func Open(filename string) (d *Deck, err error) {
	d = &Deck{}
	return
}
