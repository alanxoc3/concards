// remove ids soon
package deck

import (
	"fmt"

	"github.com/alanxoc3/concards-go/card"
	"github.com/alanxoc3/concards-go/constring"
)

type FileBreak struct {
	Id   int
	Text string
}

// DeckControl contains a deck and controls what goes in it as well as which is the current card.
type DeckControl struct {
	counter    int // incremented when cards are added, never decremented.
	Deck       Deck
	Groups     []string
	FileBreaks []FileBreak // the sections of the file that is not flash cards.
}

func (d *DeckControl) AddFileBreak(fb string) {
	d.counter += 1
	d.FileBreaks = append(d.FileBreaks, FileBreak{Id: d.counter, Text: fb})
}

// Use these to add another deck's cards to the deck.
func (d *DeckControl) AddDeckWithId(od *DeckControl) {
	for _, c := range od.Deck {
		d.AddCardWithId(c)
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

func (fb FileBreak) Print() {
	fmt.Printf("---- File Break %d Begin ----\n", fb.Id)
	fmt.Println(fb.Text)
	fmt.Printf("---- File Break %d End   ----\n", fb.Id)
}
