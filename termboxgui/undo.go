package termboxgui

import (
	"fmt"

	"github.com/alanxoc3/concards/core"
)

var stack []*core.Deck = []*core.Deck{}
var stackLocation = 0

// Saves the deck onto the change stack.
func save(d *core.Deck) {
	if len(stack) > 0 {
		// Slice is exclusive, hence the +1
		stack = stack[:stackLocation+1]
	}

	stack = append(stack, d.Copy())
	stackLocation = len(stack) - 1
}

// Returns the state of the deck, error if there are no more redo operations.
func redo() (*core.Deck, error) {
	if stackLocation+1 < len(stack) {
		stackLocation++
		d := stack[stackLocation]

		return d, nil
	} else {
		return nil, fmt.Errorf("Nothing to redo.")
	}
}

// Returns the state of the deck, error if there are no more undo operations.
func undo() (*core.Deck, error) {
	if stackLocation > 0 {
		stackLocation--
		d := stack[stackLocation]

		return d, nil
	} else {
		return nil, fmt.Errorf("Nothing to undo.")
	}
}
