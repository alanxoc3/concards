package termboxgui

import (
	"fmt"

	"github.com/alanxoc3/concards-go/card"
	"github.com/alanxoc3/concards-go/deck"
)

type stackItem struct {
	d deck.Deck
	c card.Card
}

var stack []stackItem = nil
var stack_location = 0

// Saves the deck onto the change stack.
func save(d deck.Deck) {
	si := stackItem{}
	si.d = d

	if len(d) > 0 {
		// Copy the top card.
		si.c = *(d[0])
	}

	if stack != nil {
		// Slice is exclusive, hence the +1
		stack = stack[:stack_location+1]
	}

	stack = append(stack, si)
	stack_location = len(stack) - 1
}

// Returns the state of the deck, error if there are no more redo operations.
func redo() (deck.Deck, error) {
	if stack_location < len(stack) - 1 {
		stack_location++
		d := stack[stack_location].d
		if d != nil {
			*d[0] = stack[stack_location].c
		}

		return d, nil
	} else {
		return nil, fmt.Errorf("Nothing to redo.")
	}
}

// Returns the state of the deck, error if there are no more undo operations.
func undo() (deck.Deck, error) {
	if stack_location > 0 {
		stack_location--
		d := stack[stack_location].d
		if d != nil {
			*d[0] = stack[stack_location].c
		}

		return d, nil
	} else {
		return nil, fmt.Errorf("Nothing to undo.")
	}
}
