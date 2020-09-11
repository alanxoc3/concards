package deck

import "math/rand"

func (d *Deck) Shuffle() {
	// fisher-yates shuffle
	for i := d.Len() - 1; i > 0; i-- {
		swapPlace := rand.Intn(i + 1)
		d.Swap(i, swapPlace)
	}
}
