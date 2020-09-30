package history

import (
	"fmt"

	"github.com/alanxoc3/concards/internal/deck"
)

type Manager struct {
   decks []*deck.Deck
   loc int
}

func NewManager() *Manager {
   return &Manager{[]*deck.Deck{}, 0}
}

// Saves the deck onto the change stack.
func (m *Manager) Save(decks ...*deck.Deck) {
	for _, d := range decks {
      if len(m.decks) > 0 {
         // Slice is exclusive, hence the +1
         m.decks = m.decks[:m.loc+1]
      }

      m.decks = append(m.decks, d.Copy())
      m.loc = len(m.decks) - 1
   }
}

// Returns the state of the deck, error if there are no more redo operations.
func (m *Manager) Redo() (*deck.Deck, error) {
	if m.loc+1 < len(m.decks) {
		m.loc++
		d := m.decks[m.loc]

		return d, nil
	} else {
		return nil, fmt.Errorf("Nothing to redo.")
	}
}

// Returns the state of the deck, error if there are no more undo operations.
func (m *Manager) Undo() (*deck.Deck, error) {
	if m.loc > 0 {
		m.loc--
		d := m.decks[m.loc]

		return d, nil
	} else {
		return nil, fmt.Errorf("Nothing to undo.")
	}
}
