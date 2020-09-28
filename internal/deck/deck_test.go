package deck_test

import (
	"testing"
	"time"

	"github.com/alanxoc3/concards/internal"
	"github.com/alanxoc3/concards/internal/card"
	"github.com/alanxoc3/concards/internal/deck"
	"github.com/alanxoc3/concards/internal/meta"
	"github.com/stretchr/testify/assert"
)

var ONE_DATE time.Time = time.Date(1, 1, 1, 0, 0, 0, 1, time.UTC)

func TestNewDeck(t *testing.T) {
	d := deck.NewDeck(ONE_DATE)
	assert.Len(t, d.CardList(), 0)
	assert.Len(t, d.OutcomeList(), 0)
	assert.Len(t, d.PredictList(), 0)
}

func TestAddCardsLen(t *testing.T) {
	d := deck.NewDeck(ONE_DATE)
	c1, _ := card.NewCards(".", "hi : yo")
	d.AddCards(c1...)
	assert.Len(t, d.CardList(), 2)
}

func TestAddCardsTop(t *testing.T) {
	d := deck.NewDeck(ONE_DATE)
	c1, _ := card.NewCards(".", "hi : yo")
	d.AddCards(c1...)
	assert.Equal(t, c1[0], d.TopCard())
}

func TestAddCardsPredictSameNext(t *testing.T) {
	d := deck.NewDeck(ONE_DATE)
	c1, _ := card.NewCards(".", "hi : yo")
	d.AddCards(c1...)
	p := meta.NewPredictFromStrings(c1[0].Hash().String(), "0001-01-01T00:00:00Z")
	d.AddPredicts(p)
	assert.Equal(t, 2, d.ReviewLen())
	assert.Equal(t, 0, d.FutureLen())
	assert.Equal(t, p.Hash(), *d.TopHash())
	assert.Equal(t, p, d.TopPredict())
}

func TestCardList(t *testing.T) {
	d := deck.NewDeck(ONE_DATE)
	c, _ := card.NewCards(".", "hi : yo")
	d.AddCards(c...)

	assert.Equal(t, []card.Card{*c[0], *c[1]}, d.CardList())
}

func TestPredictList(t *testing.T) {
	d := deck.NewDeck(ONE_DATE)
	h := internal.NewHash("fad")
	p := meta.NewPredictFromStrings(h.String(), "2020-01-01T00:00:00Z")
	d.AddPredicts(p)
	assert.Equal(t, []meta.Predict{*p}, d.PredictList())
}

func TestRemove(t *testing.T) {
	h1 := internal.NewHash("73a6a403534bcc11b662e3a1d90d31e1")
	h2 := internal.NewHash("abad2c2e5be5c33bc319ce038e3f2108")
	cards, _ := card.NewCards(".", "hi : yo")

	t.Run("Memorize", func(t *testing.T) {
      d := deck.NewDeck(ONE_DATE)
      d.AddCards(cards...)
		d.AddPredicts(meta.NewPredictFromStrings(h1.String()), meta.NewPredictFromStrings(h2.String(), "", "", "1"))
		d.RemoveMemorize()

		assert.Equal(t, h2, *d.TopHash())
		assert.Len(t, d.CardList(), 1)
	})

	t.Run("Review", func(t *testing.T) {
      d := deck.NewDeck(ONE_DATE)
      d.AddCards(cards...)
		d.AddPredicts(meta.NewPredictFromStrings(h1.String()), meta.NewPredictFromStrings(h2.String(), "", "", "1"))
		d.RemoveReview()

		assert.Equal(t, h1, *d.TopHash())
		assert.Len(t, d.CardList(), 1)
	})

	t.Run("Done", func(t *testing.T) {
      d := deck.NewDeck(ONE_DATE)
      d.AddCards(cards...)
      d.AddPredicts(meta.NewPredictFromStrings(h1.String()), meta.NewPredictFromStrings(h2.String(), "2020-01-01T00:00:00Z", "", "1"))
		d.RemoveDone()

		assert.Equal(t, h1, *d.TopHash())
		assert.Len(t, d.CardList(), 1)
	})
}
