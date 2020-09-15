package deck

import (
	"fmt"

	"github.com/alanxoc3/concards/internal"
	"github.com/alanxoc3/concards/internal/card"
	"github.com/alanxoc3/concards/internal/meta"
)

func removeIndex(s []internal.Hash, index int) []internal.Hash {
	return append(s[:index], s[index+1:]...)
}

// reviewStack contains checksums.
// cardMap maps checksums to cards.
// predictMap maps checksums to cards.
type Deck struct {
	reviewStack []internal.Hash                 // Hashes to review ordered by date.
	futureStack []internal.Hash                 // Hashes you have reviewed. Ordered by date.
	predictMap  map[internal.Hash]*meta.Predict // All metas.
	outcomeMap  map[internal.RKey]*meta.Outcome // Have reviewed.
	cardMap     map[internal.Hash]*card.Card    // All cards in this session.
}

func NewDeck() *Deck {
	return &Deck{
		reviewStack: []internal.Hash{},
		futureStack: []internal.Hash{},
		predictMap:  map[internal.Hash]*meta.Predict{},
		outcomeMap:  map[internal.RKey]*meta.Outcome{},
		cardMap:     map[internal.Hash]*card.Card{},
	}
}

// Deletes from the deck.
func (d *Deck) Drop(i int) error {
	if i >= 0 && i < len(d.reviewStack) {
		delete(d.cardMap, d.reviewStack[i])
		d.reviewStack = removeIndex(d.reviewStack, i)
		return nil
	} else {
		return fmt.Errorf("Can't delete. Index is out of bounds.")
	}
}

func (d *Deck) AddCard(c *card.Card) error {
	hash := c.Hash()
	_, exists := d.cardMap[hash]
	if !exists {
		d.cardMap[hash] = c
		d.reviewStack = append(d.reviewStack, hash)
		return nil
	} else {
		return fmt.Errorf("Card already exists in deck")
	}
}

func (d *Deck) InsertCard(c *card.Card, i int) error {
	hash := c.Hash()
	if i < 0 {
		i = 0
	}

	_, exists := d.cardMap[hash]
	if !exists {
		d.cardMap[hash] = c

		if i >= d.Len() {
			d.reviewStack = append(d.reviewStack, hash)
		} else {
			d.reviewStack = append(d.reviewStack, internal.Hash{})
			copy(d.reviewStack[i+1:], d.reviewStack[i:])
			d.reviewStack[i] = hash
		}

		return nil
	} else {
		return fmt.Errorf("Card already exists in deck")
	}
}

func (d *Deck) AddNewCards(file string, sides string) error {
	if cards, err := card.NewCards(file, sides); err != nil {
		return err
	} else {
		for _, c := range cards {
			if addErr := d.AddCard(c); addErr != nil {
				err = addErr
			}
		}
		return err
	}
}

func (d *Deck) AddMeta(h internal.Hash, m *meta.Predict) {
	d.predictMap[h] = m
}

func (d *Deck) AddMetaIfNil(h internal.Hash, m *meta.Predict) {
	if _, ok := d.predictMap[h]; !ok {
		d.AddMeta(h, m)
	}
}

func (d *Deck) Len() int {
	return len(d.reviewStack)
}

func (d *Deck) IsEmpty() bool {
	return d.Len() == 0
}

func (d *Deck) Swap(i, j int) {
	d.reviewStack[i], d.reviewStack[j] = d.reviewStack[j], d.reviewStack[i]
}

func (d *Deck) Get(i int) (h internal.Hash, c *card.Card, m *meta.Predict) {
	if i >= 0 && i < d.Len() {
		h = d.reviewStack[i]
		c = d.cardMap[h]
		m = d.predictMap[h]
	}
	return
}

func (d *Deck) GetHash(i int) (h internal.Hash) {
	h, _, _ = d.Get(i)
	return
}

func (d *Deck) GetCard(i int) (c *card.Card) {
	_, c, _ = d.Get(i)
	return
}

func (d *Deck) GetMeta(i int) (m *meta.Predict) {
	_, _, m = d.Get(i)
	return
}

func (d *Deck) Copy() *Deck {
	n := NewDeck()
	for _, v := range d.reviewStack {
		n.reviewStack = append(n.reviewStack, v)
	}
	for k, v := range d.cardMap {
		n.cardMap[k] = v
	}
	for k, v := range d.predictMap {
		n.predictMap[k] = v
	}
	return n
}

func (d *Deck) Clone(o *Deck) {
	d.reviewStack = []internal.Hash{}
	d.cardMap = map[internal.Hash]*card.Card{}
	d.predictMap = map[internal.Hash]*meta.Predict{}

	for _, v := range o.reviewStack {
		d.reviewStack = append(d.reviewStack, v)
	}
	for k, v := range o.cardMap {
		d.cardMap[k] = v
	}
	for k, v := range o.predictMap {
		d.predictMap[k] = v
	}
}

// Top shortcuts
func (d *Deck) Top() (internal.Hash, *card.Card, *meta.Predict) { return d.Get(0) }
func (d *Deck) TopHash() internal.Hash                          { return d.GetHash(0) }
func (d *Deck) TopCard() *card.Card                             { return d.GetCard(0) }
func (d *Deck) TopMeta() *meta.Predict                          { return d.GetMeta(0) }
func (d *Deck) DropTop() error                                  { return d.Drop(0) }

func (d *Deck) TopMetaOrDefault(defaultAlg string) *meta.Predict {
	m := d.TopMeta()
	if m == nil {
		return meta.NewDefaultPredict(d.TopHash(), defaultAlg)
	}
	return m
}

func (d *Deck) ExecTop(input bool, defaultAlg string) (*meta.Predict, error) {
	h := d.TopHash()

	if ma, e := d.TopMetaOrDefault(defaultAlg).Exec(input); e != nil {
		d.DropTop()
		return nil, e
	} else {
		d.AddMeta(h, ma)
		// Move to other deck.
		return ma, nil
	}
}

func (d *Deck) TopToEnd() {
	if len(d.reviewStack) > 1 {
		d.reviewStack = append(d.reviewStack[1:], d.reviewStack[0])
	}
}

func (d *Deck) TopTo(i int) {
	if l := len(d.reviewStack); l > 1 && i > 0 {
		if i >= l {
			i = l - 1
		}
		v := d.reviewStack[0]
		copy(d.reviewStack, d.reviewStack[1:i+1])
		d.reviewStack[i] = v
	}
}

func (d *Deck) PredictList() []meta.Predict {
	predicts := make([]meta.Predict, len(d.predictMap))
	for _, v := range d.predictMap {
		predicts = append(predicts, *v)
	}
	return predicts
}
