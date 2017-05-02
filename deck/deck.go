package deck

import (
	"fmt"

	"github.com/alanxoc3/concards-go/card"
)

// Deck contains a set of groupDecks
type Deck struct {
	counter int // incremented when cards are added, never decremented.
	Cards   []*card.Card
	Groups  []string
}

// Size returns the number of cards in the deck
func (d *Deck) Size() int {
	return len(d.Cards)
}

// Prints out the cards in the deck, for debugging purposes.
func (d *Deck) Print() {
	count := 0
	for _, c := range d.Cards {
		count += 1
		fmt.Printf("Card %d\n", count)
		c.Print()
	}
}

// Use these to add another deck's cards to the deck.
func (d *Deck) AddDeckWithId(od *Deck) {
	for _, c := range od.Cards {
		d.AddCardWithId(c)
	}
}

func (d *Deck) AddDeckWithoutId(od *Deck) {
	for _, c := range od.Cards {
		d.AddCardWithoutId(c)
	}
}

// A for some quantifier.
func isInStrList(list1 []string, item string) bool {
	for _, str := range list1 {
		if item == str {
			return true
		}
	}
	return false
}

// Treat the group splice like a set.
func (d *Deck) AddGroups(gps *[]string) {
	for _, str := range *gps {
		if !isInStrList(d.Groups, str) {
			d.Groups = append(d.Groups, str)
		}
	}
}

// These are what adds cards. The main decks should have ids.
func (d *Deck) AddCardWithoutId(c *card.Card) {
	d.counter += 1
	d.AddGroups(&c.Groups)
	d.Cards = append(d.Cards, c)
}

func (d *Deck) AddCardWithId(c *card.Card) {
	c.Id = d.counter
	d.AddCardWithoutId(c)
}
