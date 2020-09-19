package deck

import (
	"fmt"
	"time"

	"github.com/alanxoc3/concards/internal"
	"github.com/alanxoc3/concards/internal/card"
	"github.com/alanxoc3/concards/internal/meta"
)

type predictMap map[internal.Hash]*meta.Predict

// stack.review contains checksums.
// cardMap maps checksums to cards.
// predictMap maps checksums to cards.
type Deck struct {
	now         time.Time                  // Last cached time.
   stack       stack                      // Review and future stacks.
	predictMap  predictMap                 // All metas.
	outcomeMap  map[meta.Key]*meta.Outcome // Have reviewed.
	cardMap     card.CardMap               // All cards in this session.
}

func NewDeck(now time.Time) *Deck {
	return &Deck{
		now:         now,
      stack:       stack{},
		predictMap:  map[internal.Hash]*meta.Predict{},
		outcomeMap:  map[meta.Key]*meta.Outcome{},
		cardMap:     card.CardMap{},
	}
}

// Add cards only if they don't exist.
func (d *Deck) AddCards(cards ...*card.Card) {
	// See stacks.go for implementation details.
	for _, c := range cards {
		// Step 1: Cache the hash.
		h := c.Hash()

		// Step 2: Add to predict map if it isn't in there already.
		if _, exist := d.predictMap[h]; !exist {
			d.predictMap[h] = meta.NewDefaultPredict(h, "sm2")
		}

		// Step 3: Add the card only if it doesn't already exist.
		if _, exist := d.cardMap[h]; !exist {
			d.cardMap[h] = c
         d.stack.insert(h, d.predictMap, d.now)
		}
	}
}

// Add predictions only if they don't exist.
func (d *Deck) AddPredicts(predicts ...*meta.Predict) {
	for _, p := range predicts {
		h := p.Hash()
		if v, predExist := d.predictMap[h]; !predExist || v.IsZero() {
			d.predictMap[h] = p
         if _, cardExist := d.cardMap[h]; cardExist {
            d.stack.refreshHash(h, d.predictMap, d.now)
         }
      }
	}
}

func (d *Deck) ReviewLen() int { return len(d.stack.review) }
func (d *Deck) FutureLen() int { return len(d.stack.future) }

// Clones a deck into this deck.
func (d *Deck) Clone(o *Deck) {
   d.stack.clone(o.stack)
	d.cloneInfo(o)

	d.cardMap = card.CardMap{}
	for k, v := range o.cardMap {
		d.cardMap[k] = v
	}
}

func (d *Deck) Copy() *Deck {
	nd := &Deck{}
	nd.Clone(d)
	return nd
}

// Top shortcuts
func (d *Deck) TopHash() *internal.Hash {
	if len(d.stack.review) > 0 {
		return &d.stack.review[0]
	}
	return nil
}

func (d *Deck) TopCard() *card.Card {
	if len(d.stack.review) > 0 {
		return d.cardMap[d.stack.review[0]]
	}
	return nil
}

func (d *Deck) TopPredict() *meta.Predict {
	if len(d.stack.review) > 0 {
		return d.predictMap[d.stack.review[0]]
	}
	return nil
}

// Removes from both the internal map and the slice.
func (d *Deck) DropTop() {
	if len(d.stack.review) > 0 {
		delete(d.cardMap, d.stack.review[0])
		d.stack.review = d.stack.review[1:]
	}
}

// TODO: Concurrency locks for thread safety?
func (d *Deck) ExecTop(input bool, now time.Time) (meta.Predict, error) {
	// Step 1: Error if the deck is empty.
	if len(d.stack.review) == 0 {
		return *meta.NewPredictFromStrings(), fmt.Errorf("Tried to access card from an empty deck!")
	}

	// Step 2: Set the time.
	d.now = now

	// Step 3: Exec the predict value.
	p := d.TopPredict()
	np := p.Exec(input, d.now)

	// Step 4: Save the new prediction.
	d.predictMap[p.Hash()] = &np

	// Step 5: Pop the card off the review stack.
	hashToFuture := d.stack.review[0]
	d.stack.review = d.stack.review[1:]

	// Step 6: Move over some things from the future stack.
	for len(d.stack.future) > 0 && d.predictMap[d.stack.future[0]].Next().Before(d.now) {
		hashToReview := d.stack.future[0]
		d.stack.future = d.stack.future[1:]
      d.stack.insertIntoReview(hashToReview, d.predictMap)
	}

	// Step 7: Add the passed card to the future stack.
	d.stack.insertIntoFuture(hashToFuture, d.predictMap)

	return np, nil
}

// Used to write to the predict file.
func (d *Deck) PredictList() []meta.Predict {
	predicts := make([]meta.Predict, len(d.predictMap))
	for _, v := range d.predictMap {
		predicts = append(predicts, *v)
	}
	return predicts
}

// Used to write to the outcome file.
func (d *Deck) OutcomeList() []meta.Outcome {
	outcomes := make([]meta.Outcome, len(d.outcomeMap))
	for _, v := range d.outcomeMap {
		outcomes = append(outcomes, *v)
	}
	return outcomes
}

// Used for printing the cards.
func (d *Deck) CardList() []card.Card {
	cards := []card.Card{}
	for _, v := range d.stack.hashList() {
		cards = append(cards, *d.cardMap[v])
	}
	return cards
}

// Only so many cards are allowed.
func (d *Deck) Truncate(param int) {
	d.filter(func(i int, h internal.Hash) bool {
		return i < param
	})
}

func (d *Deck) RemoveMemorize() {
	d.filter(func(i int, h internal.Hash) bool {
		p := d.predictMap[h]
		return p.Total() != 0
	})
}

func (d *Deck) RemoveReview() {
	d.filter(func(i int, h internal.Hash) bool {
		p := d.predictMap[h]
		return p.Total() == 0 || p.Next().After(d.now)
	})
}

func (d *Deck) RemoveDone() {
	d.filter(func(i int, h internal.Hash) bool {
		p := d.predictMap[h]
		return p.Total() == 0 || p.Next().Before(d.now)
	})
}

type CardsFunc func(string) (card.CardMap, error)

// TODO: Concurrency locks for thread safety?
func (d *Deck) Edit(rf CardsFunc, ef CardsFunc) error {
	// Step 1: Exit if the deck is empty.
	if d.ReviewLen() == 0 {
		return fmt.Errorf("Error: The deck is empty.")
	}

	// Step 2: Cache info from the top card.
	curHash := d.TopHash()
	curCard := d.TopCard()
	curMeta := d.TopPredict()
	internal.AssertLogic(curHash != nil && curCard != nil && curMeta != nil, "no top info for non empty deck")

	filename := curCard.File()

	// Step 3: Get the current state of the file before editing it.
	beforeMap, e := rf(filename)
	if e != nil {
		return e
	}

	// Step 4: Execute the edit function.
	afterMap, e := ef(filename)
	if e != nil {
		return e
	}

	// Step 5: Remove cards that no longer exist.
	d.filter(func(i int, h internal.Hash) bool {
		_, contains := afterMap[h]
		return afterMap[h].File() != filename || contains
	})

	// Step 6: Add all the cards new after editing.
	for k, v := range afterMap {
		if _, exist := beforeMap[k]; !exist {
			if _, exist := d.predictMap[k]; !exist {
				d.AddPredicts(curMeta.Clone(v.Hash()))
			}
			d.AddCards(v)
		}
	}

	return nil
}
