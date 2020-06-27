package core

import "time"

type predicate func(int) bool

func (d *Deck) filter(p predicate) {
	for i := len(d.refs) - 1; i >= 0; i-- {
		if p(i) {
			d.Del(i)
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
		_, contains := otherDeck.Cmap[d.refs[i]]
		return d.GetCard(i).GetFile() == path && !contains
	})
}

func (d *Deck) OuterLeftJoin(otherDeck *Deck) {
	d.filter(func(i int) bool {
		_, contains := otherDeck.Cmap[d.refs[i]]
		return contains
	})
}

func (d *Deck) FilterOutFile(path string) {
	d.filter(func(i int) bool {
		return d.GetCard(i).GetFile() == path
	})
}

func (d *Deck) FilterOutMemorize() {
	d.filter(func(i int) bool {
		return d.GetMeta(i) == nil
	})
}

func (d *Deck) FilterOutReview() {
	d.filter(func(i int) bool {
		m := d.GetMeta(i)
		return m != nil && m.Next.Before(time.Now())
	})
}

func (d *Deck) FilterOutDone() {
	d.filter(func(i int) bool {
		m := d.GetMeta(i)
		return m != nil && m.Next.After(time.Now())
	})
}
