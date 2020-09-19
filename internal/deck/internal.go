package deck

import (
	"sort"
	"time"

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

func (d *Deck) filter(p predicate) {
	hashes := d.stack.hashList()
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

func beforeOrEqual(t1 time.Time, t2 time.Time) bool {
	return t1.Before(t2) || t1.Equal(t2)
}
