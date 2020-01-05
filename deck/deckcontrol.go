// remove ids soon
package deck

import (
	"fmt"

	"github.com/alanxoc3/concards/card"
	"github.com/alanxoc3/concards/constring"
)

var counter int = 0

type FileBreak struct {
	Id   int
	Text string
}

// DeckControl contains a deck and controls what goes in it as well as which is the current card.
type DeckControl struct {
	Filename   string
	Deck       Deck
	Groups     []string
	fileBreaks []FileBreak // the sections of the file that is not flash cards.
}

func (d *DeckControl) AddFileBreak(fb string) {
	counter += 1
	d.fileBreaks = append(d.fileBreaks, FileBreak{Id: counter, Text: fb})
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

func (d *DeckControl) AddGroupsSet(groups *map[string]bool) {
   for k := range *groups {
		if !constring.IsInStrList(d.Groups, k) {
			d.Groups = append(d.Groups, k)
		}
	}
}

// These are what adds cards. The main decks should have ids.
func (d *DeckControl) addCardWithoutId(c *card.Card) {
	counter += 1
	d.AddGroups(&c.Groups)
	d.Deck = append(d.Deck, c)
}

func (d *DeckControl) AddCardWithId(c *card.Card) {
	c.Id = counter
	d.addCardWithoutId(c)
}

// returns a string of all the cards in their sorted order.
func (d *DeckControl) ToString() string {
	// Step 1: Make a sorted copy of the deck.
	list := make(Deck, len(d.Deck))
	copy(list, d.Deck)
	list.Sort()

	// Step 2: Make a string and temporary deck.
	str := ""
	var tmp Deck

	// Step 3: Loop through both File Breaks and Cards, putting the ids in order.
	// i is for list, j is for d.fileBreaks
	for i, j := 0, 0; i+j <= len(list)+len(d.fileBreaks); {
		listFin := (i == len(list))
		breakFin := (j == len(d.fileBreaks))
		done := (breakFin && listFin)
		listTurn := !breakFin && !listFin && (list[i].Id < d.fileBreaks[j].Id) || breakFin && !listFin
		breakTurn := !listTurn

		if done || breakTurn {
			if len(tmp) != 0 {
				str += tmp.ToString()
				str += "##\n"
				tmp = nil
			}
		}

		if done {
			break
		} else if listTurn {
         if !list[i].Deleted {
            tmp = append(tmp, list[i])
         }
			i++
		} else if breakTurn {
			str += d.fileBreaks[j].Text
			j++
		}

		//fmt.Printf("i is %d, j is %d\n", i, j)
	}
	return str
}

func (fb FileBreak) Print() {
	fmt.Printf("---- File Break %d Begin ----\n", fb.Id)
	fmt.Println(fb.Text)
	fmt.Printf("---- File Break %d End   ----\n", fb.Id)
}
