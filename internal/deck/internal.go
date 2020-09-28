package deck

import (
	"time"

	"github.com/alanxoc3/concards/internal"
	"github.com/alanxoc3/concards/internal/card"
	"github.com/alanxoc3/concards/internal/meta"
)

type predicate func(int, internal.Hash) bool

func (d *Deck) cloneInfo(o *Deck) {
	d.stack.SetTime(o.stack.Time())
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
	hashes := d.stack.List()
	n := &Deck{}
	n.cloneInfo(d)

	for i, h := range hashes {
		if p(i, h) {
			n.AddCards(d.cardMap[h])
		}
	}
	d.Clone(n)
}

func beforeOrEqual(t1 time.Time, t2 time.Time) bool {
	return t1.Before(t2) || t1.Equal(t2)
}

func cardListToMap(cl []*card.Card) card.CardMap {
	cm := card.CardMap{}
	for _, c := range cl {
		h := c.Hash()
		if _, exist := cm[h]; !exist {
			cm[h] = c
		}
	}
	return cm
}
