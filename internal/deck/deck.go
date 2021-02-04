package deck

import (
	"fmt"
	"time"

	"github.com/alanxoc3/concards/internal"
	"github.com/alanxoc3/concards/internal/card"
	"github.com/alanxoc3/concards/internal/meta"
	"github.com/alanxoc3/concards/internal/stack"
)

// stack.review contains checksums.
// cardMap maps checksums to cards.
// predictMap maps checksums to cards.
type Deck struct {
	stack      stack.Stack     // Review and future stacks.
	cardMap    card.CardMap    // All cards.
	predictMap meta.PredictMap // All metas.
	outcomeMap meta.OutcomeMap // Have reviewed.
}

func NewDeck(now time.Time) *Deck {
	return &Deck{
		stack:      stack.NewStack(now),
      cardMap:    map[internal.Hash]*card.Card{},
		predictMap: map[internal.Hash]*meta.Predict{},
		outcomeMap: map[meta.Key]*meta.Outcome{},
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

		// Step 3: Add to the stack.
		if _, exist := d.cardMap[h]; !exist {
			d.cardMap[h] = c
		}

      d.stack.Insert(h, d.predictMap[h].Next(), d.predictMap[h].Total() == 0)
	}
}

// Add predictions only if they don't exist.
func (d *Deck) AddPredicts(predicts ...*meta.Predict) {
	for _, p := range predicts {
		h := p.Hash()
		if v, predExist := d.predictMap[h]; !predExist || v.IsZero() {
			d.predictMap[h] = p
			d.stack.Update(h, p.Next(), p.Total() == 0)
		}
	}
}

func (d *Deck) ReviewLen() int { return d.stack.ReviewLen() }
func (d *Deck) FutureLen() int { return d.stack.FutureLen() }

// Clones a deck into this deck.
func (d *Deck) Clone(o *Deck) {
	d.stack.Clone(o.stack)

	d.cardMap = map[internal.Hash]*card.Card{}
	for k, v := range o.cardMap {
		d.cardMap[k] = v
	}

	d.cloneInfo(o)
}

func (d *Deck) Copy() *Deck {
	nd := &Deck{}
	nd.Clone(d)
	return nd
}

// Top shortcuts
func (d *Deck) TopHash() *internal.Hash {
   return d.stack.Top()
}

func (d *Deck) TopCard() *card.Card {
   if h := d.TopHash(); h != nil {
		return d.cardMap[*h]
   }
	return nil
}

func (d *Deck) TopPredict() *meta.Predict {
   if h := d.TopHash(); h != nil {
		return d.predictMap[*h]
   }
	return nil
}

// Removes from both the internal map and the slice.
func (d *Deck) DropTop() {
   if h := d.TopHash(); h != nil {
		delete(d.cardMap, *h)
      d.stack.Pop()
	}
}

// TODO: Concurrency locks for thread safety?
func (d *Deck) ExecTop(input bool, now time.Time) (meta.Predict, error) {
	// Step 1: Error if the deck is empty.
	if d.stack.ReviewLen() == 0 {
		return *meta.NewPredictFromStrings(), fmt.Errorf("Tried to access card from an empty deck!")
	}

	// Step 2: Exec the predict value.
	np := d.TopPredict().Exec(input, now)
   no := meta.NewOutcomeFromPredict(&np, now, input)

	// Step 3: Save the new prediction.
	d.predictMap[np.Hash()] = &np

	// Step 4: Save the outcome too.
   d.outcomeMap[no.Key()] = no

   // Step 5: Set the current time.
   d.stack.SetTime(now)

   // Step 6: Update the stack.
   updateStatus := d.stack.Update(np.Hash(), np.Next(), np.Total() == 0)
   internal.AssertLogic(updateStatus, "stack didn't contain hash")

   return np, nil
}

// Used to write to the predict file.
func (d *Deck) PredictList() []*meta.Predict {
	predicts := []*meta.Predict{}
	for _, v := range d.predictMap {
		predicts = append(predicts, v)
	}
	return predicts
}

// Used to write to the outcome file.
func (d *Deck) OutcomeList() []*meta.Outcome {
	outcomes := []*meta.Outcome{}
	for _, v := range d.outcomeMap {
		outcomes = append(outcomes, v)
	}
	return outcomes
}

// Used for printing the cards.
func (d *Deck) CardList() []*card.Card {
	cards := []*card.Card{}
	for _, v := range d.stack.List() {
		cards = append(cards, d.cardMap[v])
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
		return p.Total() == 0 || p.Next().After(d.stack.Time())
	})
}

func (d *Deck) RemoveDone() {
	d.filter(func(i int, h internal.Hash) bool {
		p := d.predictMap[h]
		return p.Total() == 0 || beforeOrEqual(p.Next(), d.stack.Time())
	})
}

type CardsFunc func(string) ([]*card.Card, error)

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
	beforeList, e := rf(filename)
	if e != nil { return e }

	// Step 4: Execute the edit function.
	afterList, e := ef(filename)
	if e != nil {
		return e
	}

	// Step 5: Remove cards that no longer exist.
   afterMap := cardListToMap(afterList)
	d.filter(func(i int, h internal.Hash) bool {
		_, contains := afterMap[h]
		return d.cardMap[h].File() != filename || contains
	})

	// Step 6: Add all the cards new after editing.
   beforeMap := cardListToMap(beforeList)
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
