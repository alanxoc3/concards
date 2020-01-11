package deck

import (
	"math/rand"
	"sort"

	"github.com/alanxoc3/concards/card"
)

type Deck []*card.Card

func (cards Deck) Len() int {
	return len(cards)
}

func (cards Deck) Less(i, j int) bool {
   return cards[i].Question < cards[j].Question
}

func (cards Deck) Swap(i, j int) {
	cards[i], cards[j] = cards[j], cards[i]
}

func (cards Deck) Sort() {
	sort.Sort(cards)
}

// fisher-yates shuffle
func (cards Deck) Shuffle() {
	// start at the end of the deck, go down.
	for i := len(cards) - 1; i > 0; i-- {
		swapPlace := rand.Intn(i + 1) // The plus one is to enable the card to remain in the same place.
		cards.Swap(i, swapPlace)
	}
}

func (cards Deck) Top() *card.Card {
	if len(cards) > 0 {
		return cards[0]
	} else {
		return nil
	}
}
