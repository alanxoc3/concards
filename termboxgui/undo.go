package termboxgui

import (
	"fmt"

	"github.com/alanxoc3/concards/core"
)

var stack []*core.Deck = []*core.Deck{}
var stack_location = 0

// Saves the deck onto the change stack.
func save(d *core.Deck) {
	if len(stack) > 0 {
		// Slice is exclusive, hence the +1
		stack = stack[:stack_location+1]
	}

	stack = append(stack, d.Copy())
	stack_location = len(stack) - 1
}

// Returns the state of the deck, error if there are no more redo operations.
func redo() (*core.Deck, error) {
	if stack_location < len(stack) - 1 {
		stack_location++
		d := stack[stack_location]

		return d, nil
	} else {
		return nil, fmt.Errorf("Nothing to redo.")
	}
}

// Returns the state of the deck, error if there are no more undo operations.
func undo() (*core.Deck, error) {
	if stack_location > 0 {
		stack_location--
		d := stack[stack_location]

		return d, nil
	} else {
		return nil, fmt.Errorf("Nothing to undo.")
	}
}
