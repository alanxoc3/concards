package deck

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/alanxoc3/concards-go/card"
	"github.com/alanxoc3/concards-go/constring"
)

type Deck []*card.Card

func (cards Deck) Len() int {
	return len(cards)
}

func (cards Deck) Less(i, j int) bool {
	return cards[i].Id < cards[j].Id
}

func (cards Deck) Swap(i, j int) {
	cards[i], cards[j] = cards[j], cards[i]
}

func (cards Deck) Sort() {
	sort.Sort(cards)
}

// includes the group markers
func (cards Deck) ToString() string {
	// do groups stuff
	str := ""
	var curGroups []string

	for _, c := range cards {
		if !constring.StringListsIdentical(curGroups, c.Groups) {
			curGroups = c.Groups
			str += constring.GroupListToString(curGroups) + "\n"
		}

		str += c.FormatFile() + "\n\n"
	}

	return str
}

// fisher-yates shuffle
func (cards Deck) Shuffle() {
	// start at the end of the deck, go down.
	for i := len(cards) - 1; i > 0; i-- {
		swapPlace := rand.Intn(i)
		cards.Swap(i, swapPlace)
	}
}

// Prints out the cards in the deck, for debugging purposes.
func (cards Deck) Print() {
	count := 0
	for _, c := range cards {
		count += 1
		fmt.Printf("Card %d\n", count)
		c.Print()
	}
}

func (cards Deck) Top() *card.Card {
	if len(cards) > 0 {
		return cards[0]
	} else {
		return nil
	}
}
