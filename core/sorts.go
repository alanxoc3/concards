package core

import (
	"math/rand"
	"sort"
)

func (d *Deck) Less(i, j int) bool {
	return d.Cmap[d.refs[i]].GetQuestion() < d.Cmap[d.refs[j]].GetQuestion()
}

func (d *Deck) SortByQuestion() {
	sort.Sort(d)
}

func (d *Deck) Shuffle() {
	// fisher-yates shuffle
	for i := d.Len() - 1; i > 0; i-- {
		swapPlace := rand.Intn(i + 1)
		d.Swap(i, swapPlace)
	}
}
