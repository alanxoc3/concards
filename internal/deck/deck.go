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
func (d *Deck) drop(i int) error {
	if i >= 0 && i < len(d.reviewStack) {
		delete(d.cardMap, d.reviewStack[i])
		d.reviewStack = removeIndex(d.reviewStack, i)
		return nil
	} else {
		return fmt.Errorf("Can't delete. Index is out of bounds.")
	}
}

func (d *Deck) addCard(c *card.Card) error {
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

		if i >= d.ReviewLen() {
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

func (d *Deck) AddCards(cards ...*card.Card) {
   for _, c := range cards {
      hash := c.Hash()
      if _, exist := d.cardMap[hash]; !exist {
         d.cardMap[hash] = c
         d.reviewStack = append(d.reviewStack, hash)
      }
   }
}

func (d *Deck) AddPredicts(predicts ...*meta.Predict) {
   for _, p := range predicts {
      h := p.Hash()
      if _, exist := d.predictMap[h]; !exist {
         d.predictMap[p.Hash()] = p
      }
   }
}

func (d *Deck) ReviewLen() int { return len(d.reviewStack) }
func (d *Deck) FutureLen() int { return len(d.futureStack) }

func (d *Deck) getHash(i int) *internal.Hash {
	internal.AssertLogic(i >= 0 && i < d.ReviewLen(), "index out of bounds")
	return &d.reviewStack[i]
}

func (d *Deck) getCard(i int) *card.Card {
	internal.AssertLogic(i >= 0 && i < d.ReviewLen(), "index out of bounds")
	return d.cardMap[d.reviewStack[i]]
}

func (d *Deck) getPredict(i int) *meta.Predict {
	internal.AssertLogic(i >= 0 && i < d.ReviewLen(), "index out of bounds")
	return d.predictMap[d.reviewStack[i]]
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
func (d *Deck) TopHash() *internal.Hash {
	if d.ReviewLen() == 0 { return nil
   } else { return d.getHash(0) }
}

func (d *Deck) TopCard() *card.Card {
	if d.ReviewLen() == 0 { return nil
   } else { return d.getCard(0) }
}

func (d *Deck) TopPredict() *meta.Predict {
	if d.ReviewLen() == 0 { return nil
   } else { return d.getPredict(0) }
}

// TODO: don't depend on the drop method.
func (d *Deck) DropTop() error { return d.drop(0) }

func (d *Deck) ExecTop(input bool, defaultAlg string) (meta.Predict, error) {
   if d.ReviewLen() == 0 { return *meta.NewPredictFromStrings(), fmt.Errorf("Tried to access card from an empty deck!") }
   p := d.TopPredict()

   // TODO: Move the card & create a meta outcome.
	np := p.Exec(input);
   d.predictMap[p.Hash()] = &np
   return np, nil
}

func (d *Deck) PredictList() []meta.Predict {
	predicts := make([]meta.Predict, len(d.predictMap))
	for _, v := range d.predictMap {
		predicts = append(predicts, *v)
	}
	return predicts
}

func (d *Deck) OutcomeList() []meta.Outcome {
	outcomes := make([]meta.Outcome, len(d.outcomeMap))
	for _, v := range d.outcomeMap {
		outcomes = append(outcomes, *v)
	}
	return outcomes
}

func (d *Deck) CardList() []card.Card {
	cards := make([]card.Card, len(d.cardMap))
	for _, v := range d.cardMap {
		cards = append(cards, *v)
	}
	return cards
}
