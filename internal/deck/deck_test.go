package deck_test

import (
	"testing"
	"time"

	"github.com/alanxoc3/concards/internal/card"
	"github.com/alanxoc3/concards/internal/deck"
	"github.com/alanxoc3/concards/internal/meta"
	"github.com/stretchr/testify/assert"
)

var ONE_DATE  time.Time = time.Date(1, 1, 1, 0, 0, 0, 1, time.UTC)

func TestNewDeck(t *testing.T) {
	d := deck.NewDeck(ONE_DATE)
	assert.Len(t, d.CardList(), 0)
	assert.Len(t, d.OutcomeList(), 0)
	assert.Len(t, d.PredictList(), 0)
}

func TestAddCardsTop(t *testing.T) {
	d := deck.NewDeck(ONE_DATE)
	c1, _ := card.NewCards(".", "hi : yo")
	d.AddCards(c1...)
	assert.Len(t, d.CardList(), 2)
	assert.Equal(t, c1[0], d.TopCard())
}

func TestAddCardsPredictSameNext(t *testing.T) {
	d := deck.NewDeck(ONE_DATE)
	c1, _ := card.NewCards(".", "hi : yo")
	d.AddCards(c1...)
   p := meta.NewPredictFromStrings(c1[0].Hash().String(), "", "2020-01-01T00:00:00Z")
	d.AddPredicts(p)
	assert.Equal(t, 2, d.ReviewLen())
	assert.Equal(t, p.Hash(), *d.TopHash())
	assert.Equal(t, p, d.TopPredict())
}
