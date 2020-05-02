package con

import (
   "math/rand"
   "sort"
)

func (d Deck) Less(i, j int) bool {
   return d.cmap[d.refs[i]].GetQuestion() < d.cmap[d.refs[j]].GetQuestion()
}

func (d Deck) SortByQuestion() {
   sort.Sort(d)
}

// fisher-yates shuffle
func (d Deck) Shuffle() {
   // start at the end of the deck, go down.
   for i := d.Len() - 1; i > 0; i-- {
      swapPlace := rand.Intn(i + 1) // The plus one is to enable the card to remain in the same place.
      d.Swap(i, swapPlace)
   }
}
