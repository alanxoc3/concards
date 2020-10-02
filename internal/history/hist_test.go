package history_test

import (
	"testing"
	"time"

	"github.com/alanxoc3/concards/internal/card"
	"github.com/alanxoc3/concards/internal/deck"
	"github.com/alanxoc3/concards/internal/history"
	"github.com/stretchr/testify/assert"
)

var ONE_DATE time.Time = time.Date(1, 1, 1, 0, 0, 0, 1, time.UTC)

func create(strs ...string) *history.Manager {
   m := history.NewManager()
	for _, s := range strs {
      d := deck.NewDeck(ONE_DATE)
      c, _ := card.NewCards(".", s)
      d.AddCards(c...)
      m.Save(d)
   }
	return m
}

func TestUndo(t *testing.T) {
   m := create("1", "2", "3")
	d, e := m.Undo()
	assert.Equal(t, "2", d.TopCard().GetFactEsc(0))
	assert.Nil(t, e)
}

func TestRedo(t *testing.T) {
   m := create("1", "2", "3")
	m.Undo()
	d, e := m.Redo()
	assert.Equal(t, "3", d.TopCard().GetFactEsc(0))
	assert.Nil(t, e)
}

func TestUndoEnd(t *testing.T) {
   m := create("1")
	d, e := m.Undo()
	assert.Nil(t, d)
	assert.NotNil(t, e)
}

func TestRedoEnd(t *testing.T) {
   m := create("1")
	d, e := m.Redo()
	assert.Nil(t, d)
	assert.NotNil(t, e)
}

func TestUndoEmpty(t *testing.T) {
   m := create()
	d, e := m.Undo()
	assert.Nil(t, d)
	assert.NotNil(t, e)
}

func TestRedoEmpty(t *testing.T) {
   m := create()
	d, e := m.Redo()
	assert.Nil(t, d)
	assert.NotNil(t, e)
}
