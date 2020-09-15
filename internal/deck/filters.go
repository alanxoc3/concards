package deck

import "time"

type predicate func(int) bool

func (d *Deck) filter(p predicate) {
	for i := len(d.reviewStack) - 1; i >= 0; i-- {
		if p(i) {
			d.drop(i)
		}
	}
}

// Basically truncates a deck.
func (d *Deck) FilterNumber(param int) {
	d.filter(func(i int) bool {
		return i >= param
	})
}

func (d *Deck) FileIntersection(path string, otherDeck *Deck) {
	d.filter(func(i int) bool {
		_, contains := otherDeck.predictMap[d.reviewStack[i]]
		return d.getCard(i).File() == path && !contains
	})
}

func (d *Deck) OuterLeftJoin(otherDeck *Deck) {
	d.filter(func(i int) bool {
		_, contains := otherDeck.predictMap[d.reviewStack[i]]
		return contains
	})
}

func (d *Deck) FilterOutFile(path string) {
	d.filter(func(i int) bool {
		return d.getCard(i).File() == path
	})
}

func (d *Deck) FilterOutMemorize() {
	d.filter(func(i int) bool {
      p := d.getPredict(i)
		return p == nil || p.Total() == 0
	})
}

func (d *Deck) FilterOutReview() {
   now := time.Now()
	d.filter(func(i int) bool {
      p := d.getPredict(i)
		return p != nil && p.Next().Before(now)
	})
}

func (d *Deck) FilterOutDone() {
   now := time.Now()
	d.filter(func(i int) bool {
      p := d.getPredict(i)
		return p != nil && p.Next().After(now)
	})
}
