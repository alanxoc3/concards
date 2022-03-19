package deck_test

import (
	"testing"
	"time"

	"github.com/alanxoc3/concards/internal"
	"github.com/alanxoc3/concards/internal/card"
	"github.com/alanxoc3/concards/internal/deck"
	"github.com/alanxoc3/concards/internal/meta"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var DATE_0 time.Time = time.Time{}
var DATE_1 time.Time = time.Date(1, 1, 1, 0, 0, 1, 0, time.UTC)
var DATE_2 time.Time = time.Date(1, 1, 1, 0, 0, 2, 0, time.UTC)
var DATE_3 time.Time = time.Date(1, 1, 1, 0, 0, 3, 0, time.UTC)
var DATE_4 time.Time = time.Date(1, 1, 1, 0, 0, 4, 0, time.UTC)
var DATE_5 time.Time = time.Date(1, 1, 1, 0, 0, 5, 0, time.UTC)

func TestNewDeck(t *testing.T) {
	d := deck.NewDeck(DATE_1)
	assert.Len(t, d.CardList(), 0)
	assert.Len(t, d.OutcomeList(), 0)
	assert.Len(t, d.PredictList(), 0)
	assert.Nil(t, d.TopCard())
	assert.Nil(t, d.TopPredict())
	assert.Nil(t, d.TopHash())
}

func TestAddCardsLen(t *testing.T) {
	d := deck.NewDeck(DATE_1)
	cards, _ := card.NewCards(".", "hi :: yo")
	d.UpsertCards(cards...)
	assert.Len(t, d.CardList(), 2)
}

func TestAddCardsTop(t *testing.T) {
	d := deck.NewDeck(DATE_1)
	cards, _ := card.NewCards(".", "hi :: yo")
	d.UpsertCards(cards...)
	assert.Equal(t, cards[0], d.TopCard())
}

func TestDropTop(t *testing.T) {
	d := deck.NewDeck(DATE_1)
	cards, _ := card.NewCards(".", "hi :: yo")
	d.UpsertCards(cards...)
	d.DropTop()
	assert.Equal(t, cards[1], d.TopCard())
	assert.Len(t, d.CardList(), 1)
}

func TestTruncate(t *testing.T) {
	d := deck.NewDeck(DATE_1)
	cards, _ := card.NewCards(".", "hi :: yo")
	d.UpsertCards(cards...)
	d.Truncate(1)
	assert.Equal(t, cards[0], d.TopCard())
	assert.Len(t, d.CardList(), 1)
}

func TestHardCardsAreFirstWhenAddingPredictsThenCards(t *testing.T) {
	d := deck.NewDeck(DATE_3)
	cards, _ := card.NewCards(".", "hi :: yo")
	p := meta.NewPredictFromStrings(cards[0].Hash().String(), "0001-01-01T00:00:02Z", "0001-01-01T00:00:01Z", "1")
	d.UpsertPredicts(p)
	d.UpsertCards(cards...)
	assert.Equal(t, cards[1], d.TopCard())
	assert.Len(t, d.CardList(), 2)
}

func TestHardCardsAreFirstWhenAddingCardsThenPredicts(t *testing.T) {
	d := deck.NewDeck(DATE_3)
	cards, _ := card.NewCards(".", "hi :: yo")
	d.UpsertCards(cards...)
	require.Equal(t, cards[0], d.TopCard())
	require.Len(t, d.CardList(), 2)

	p := meta.NewPredictFromStrings(cards[0].Hash().String(), "0001-01-01T00:00:02Z", "0001-01-01T00:00:01Z", "1")
	d.UpsertPredicts(p)
	assert.Equal(t, cards[1], d.TopCard())
	assert.Len(t, d.CardList(), 2)
}

func TestHardCardsAreFirstWhenUsingExecTop(t *testing.T) {
	d := deck.NewDeck(DATE_1)
	cards, _ := card.NewCards(".", "hi :: yo | me")

	p1 := meta.NewPredictFromStrings(cards[0].Hash().String(), "0001-01-01T00:00:01Z", "0001-01-01T00:00:01Z", "1", "0", "0", "unit-test")
	p2 := meta.NewPredictFromStrings(cards[1].Hash().String(), "", "", "", "", "", "unit-test")
	p3 := meta.NewPredictFromStrings(cards[2].Hash().String(), "", "", "", "", "", "unit-test")
	d.UpsertPredicts(p1, p2, p3)
	d.UpsertCards(cards...)

	require.Equal(t, 2, d.Len())
	require.Equal(t, cards[1], d.TopCard())
	d.ExecTop(false, DATE_1)

	require.Equal(t, 1, d.Len())
	require.Equal(t, cards[2], d.TopCard())
	d.ExecTop(false, DATE_2)

	require.Equal(t, 2, d.Len())
	require.Equal(t, cards[0], d.TopCard())
	d.ExecTop(true, DATE_3)

	require.Equal(t, 2, d.Len())
	require.Equal(t, cards[2], d.TopCard())
	d.ExecTop(true, DATE_4)

	require.Equal(t, 1, d.Len())
	require.Equal(t, cards[1], d.TopCard())
	d.ExecTop(true, DATE_5)

	require.Zero(t, d.Len())
	require.Nil(t, d.TopCard())
}

func TestAddCardsPredictSameNext(t *testing.T) {
	d := deck.NewDeck(DATE_1)
	cards, _ := card.NewCards(".", "hi :: yo")
	d.UpsertCards(cards...)
	p := meta.NewPredictFromStrings(cards[0].Hash().String(), "0001-01-01T00:00:00Z")
	d.UpsertPredicts(p)
	assert.Equal(t, 2, d.Len())
	assert.Equal(t, 2, d.Capacity())
	// assert.Equal(t, p.Hash(), *d.TopHash())
	// assert.Equal(t, p, d.TopPredict())
}

func TestCardList(t *testing.T) {
	d := deck.NewDeck(DATE_1)
	c, _ := card.NewCards(".", "hi :: yo")
	d.UpsertCards(c...)

	assert.Equal(t, c, d.CardList())
}

func TestPredictList(t *testing.T) {
	d := deck.NewDeck(DATE_1)
	h := internal.NewHash("fad")
	p := meta.NewPredictFromStrings(h.String(), "2020-01-01T00:00:00Z")
	d.UpsertPredicts(p)
	assert.Equal(t, []*meta.Predict{p}, d.PredictList())
}

func TestRemove(t *testing.T) {
	h1 := internal.NewHash("73a6a403534bcc11b662e3a1d90d31e1")
	h2 := internal.NewHash("abad2c2e5be5c33bc319ce038e3f2108")
	cards, _ := card.NewCards(".", "hi :: yo")

	t.Run("Memorize", func(t *testing.T) {
		d := deck.NewDeck(DATE_1)
		d.UpsertCards(cards...)
		d.UpsertPredicts(meta.NewPredictFromStrings(h1.String()), meta.NewPredictFromStrings(h2.String(), "", "", "1"))
		d.RemoveMemorize()

		assert.Equal(t, h2, *d.TopHash())
		assert.Len(t, d.CardList(), 1)
	})

	t.Run("Review", func(t *testing.T) {
		d := deck.NewDeck(DATE_1)
		d.UpsertCards(cards...)
		d.UpsertPredicts(meta.NewPredictFromStrings(h1.String()), meta.NewPredictFromStrings(h2.String(), "", "", "1"))
		d.RemoveReview()

		assert.Equal(t, h1, *d.TopHash())
		assert.Len(t, d.CardList(), 1)
	})

	t.Run("Done", func(t *testing.T) {
		d := deck.NewDeck(DATE_1)
		d.UpsertCards(cards...)
		d.UpsertPredicts(meta.NewPredictFromStrings(h1.String()), meta.NewPredictFromStrings(h2.String(), "2020-01-01T00:00:00Z", "", "1"))
		d.RemoveDone()

		assert.Equal(t, h1, *d.TopHash())
		assert.Len(t, d.CardList(), 1)
	})
}

func TestExecFuture(t *testing.T) {
	h1 := internal.NewHash("55746153a816f94836725a939a5cac37")
	h2 := internal.NewHash("abad2c2e5be5c33bc319ce038e3f2108")
	h3 := internal.NewHash("a64332565db9aebe4f2edc4f1c125610")
	cards, _ := card.NewCards(".", "hi :: yo | me")

	d := deck.NewDeck(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
	d.UpsertCards(cards...)
	d.UpsertPredicts(
		meta.NewPredictFromStrings(h1.String(), "2020-01-02T00:00:00Z", "2020-01-02T00:00:00Z", "1"),
		meta.NewPredictFromStrings(h2.String()),
		meta.NewPredictFromStrings(h3.String()),
	)

	assert.Equal(t, 1, d.Capacity() - d.Len())
	assert.Equal(t, h2, *d.TopHash())
	d.ExecTop(false, time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC))

	assert.Len(t, d.OutcomeList(), 1)
	assert.Equal(t, h3, *d.TopHash())
	assert.Equal(t, 1, d.Capacity() - d.Len())
}

func TestEdit(t *testing.T) {
	commonDeck := deck.NewDeck(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
	cards, _ := card.NewCards(".", "hi | yo")
	commonDeck.UpsertCards(cards...)

	t.Run("New", func(t *testing.T) {
		d := commonDeck.Copy()
		d.Edit(func(s string) ([]*card.Card, error) {
			return card.NewCards(".", "hi :: yo")
		})

		assert.Equal(t, 2, d.Len())
		assert.Equal(t, cards[0], d.TopCard())
	})

	t.Run("Replace", func(t *testing.T) {
		d := commonDeck.Copy()
		c, _ := card.NewCards(".", "yo | hi")

		d.Edit(func(s string) ([]*card.Card, error) {
			return c, nil
		})

		assert.Equal(t, 1, d.Len())
		assert.Equal(t, c[0], d.TopCard())
	})

	t.Run("Delete", func(t *testing.T) {
		d := commonDeck.Copy()
		d.Edit(func(s string) ([]*card.Card, error) {
			return []*card.Card{}, nil
		})

		assert.Equal(t, 0, d.Len())
	})
}
