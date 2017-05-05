package deck

import (
	"github.com/alanxoc3/concards-go/card"
	"github.com/alanxoc3/concards-go/constring"
)

// DeckControl contains a deck and controls what goes in it as well as which is the current card.
type DeckControl struct {
	counter int // incremented when cards are added, never decremented.
	Deck    Deck
	Groups  []string
}

// Use these to add another deck's cards to the deck.
func (d *DeckControl) AddDeckWithId(od *DeckControl) {
	for _, c := range od.Deck {
		d.AddCardWithId(c)
	}
}

func (d *DeckControl) AddDeckWithoutId(od *DeckControl) {
	for _, c := range od.Deck {
		d.AddCardWithoutId(c)
	}
}

// Treat the group splice like a set.
func (d *DeckControl) AddGroups(gps *[]string) {
	for _, str := range *gps {
		if !constring.IsInStrList(d.Groups, str) {
			d.Groups = append(d.Groups, str)
		}
	}
}

// These are what adds cards. The main decks should have ids.
func (d *DeckControl) AddCardWithoutId(c *card.Card) {
	d.counter += 1
	d.AddGroups(&c.Groups)
	d.Deck = append(d.Deck, c)
}

func (d *DeckControl) AddCardWithId(c *card.Card) {
	c.Id = d.counter
	d.AddCardWithoutId(c)
}

// given a file name, returns a string of all the cards part of that file.
func (d *DeckControl) ToStringFromFile(file string) string {
	var list Deck
	list = d.Deck.FilterFile(file)
	list.Sort()
	return list.ToString()
}

// given a bunch of groups,
func (d *DeckControl) ToStringFromGroups(groups []string) string {
	var list Deck
	list = d.Deck.FilterGroups(groups)
	list.Sort()
	return list.ToString()
}

// given a file name, returns a string of all the cards part of that file.
func (d *DeckControl) ToStringFromGroup(group string) string {
	var list Deck
	list = d.Deck.FilterGroup(group)
	list.Sort()
	return list.ToString()
}
