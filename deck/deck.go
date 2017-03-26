package deck

import (
	"fmt"

	"github.com/alanxoc3/concards-go/card"
)

// Contains the group divisions, as parsed in the file.
type SubDeck struct {
	Cards  []*card.Card
	Groups []string
}

// Deck contains a set of groupDecks
type Deck struct {
	SubDecks []*SubDeck
}

// Size for a subdeck
func (d *SubDeck) Size() int {
	return len(d.Cards)
}

// Size returns the number of cards in the deck
func (d *Deck) Size() (total int) {
	for _, x := range d.SubDecks {
		total += x.Size()
	}
	return
}

// Prints out the cards in the deck, for debugging purposes.
func (d *Deck) Print() {
	count := 0
	for _, sd := range d.SubDecks {
		for _, c := range sd.Cards {
			count += 1
			fmt.Printf("Card %d\n", count)
			c.Print()
		}
	}
}

// Adds the card, along with setting the card's group.
func (d *SubDeck) AddCard(c *card.Card) {
	c.Groups = d.Groups
	d.Cards = append(d.Cards, c)
}
