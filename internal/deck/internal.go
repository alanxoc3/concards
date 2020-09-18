package deck

import (
	"sort"

	"github.com/alanxoc3/concards/internal"
	"github.com/alanxoc3/concards/internal/meta"
)

type predicate func(int, internal.Hash) bool

func (d *Deck) cloneInfo(o *Deck) {
	d.now = o.now
	d.predictMap = map[internal.Hash]*meta.Predict{}
	for k, v := range o.predictMap {
		d.predictMap[k] = v
	}

	d.outcomeMap = map[meta.Key]*meta.Outcome{}
	for k, v := range o.outcomeMap {
		d.outcomeMap[k] = v
	}
}

func (d *Deck) hashList() []internal.Hash {
	hashes := []internal.Hash{}
	hashes = append(hashes, d.reviewStack...)
	hashes = append(hashes, d.futureStack...)
	return hashes
}

func (d *Deck) filter(p predicate) {
	hashes := d.hashList()
	nd := &Deck{}
	nd.cloneInfo(d)

	for i, h := range hashes {
		if p(i, h) {
			nd.AddCards(d.cardMap[h])
		}
	}
	d.Clone(nd)
}

func insertSorted(hs []internal.Hash, h internal.Hash, lessFunc func(int) bool) []internal.Hash {
	i := sort.Search(len(hs), lessFunc)
	hs = append(hs, internal.Hash{})
	copy(hs[i+1:], hs[i:])
	hs[i] = h
	return hs
}

// Requires the card to have an entry in both the predictMap & cardMap.
func (d *Deck) insertIntoReviewStack(h internal.Hash) {
	next := d.predictMap[h].Next()
	d.reviewStack = insertSorted(d.reviewStack, h, func(i int) bool {
		return d.predictMap[d.reviewStack[i]].Next().Before(next)
	})
}

// Requires the card to have an entry in both the predictMap & cardMap.
func (d *Deck) insertIntoFutureStack(h internal.Hash) {
	next := d.predictMap[h].Next()
	d.futureStack = insertSorted(d.futureStack, h, func(i int) bool {
		return d.predictMap[d.futureStack[i]].Next().After(next)
	})
}

// Returns a new stack with the hash removed.
func removeFromStack(stack []internal.Hash, h internal.Hash) []internal.Hash {
	// TODO: Make this more effecient?
	newStack := []internal.Hash{}
	for _, v := range stack {
		if v != h {
			newStack = append(newStack, v)
		}
	}

	if len(newStack) < len(stack) {
		return newStack
	}
	return stack
}
