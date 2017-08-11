package termboxgui

const (
	Edit = iota
	No = iota
	Idk = iota
	Yes = iota
)

// Undo is a stack. "undo" pops off the stack.
// Redo pushes onto the stack.

var stack []deck.Deck
var stack_location := 0

func undo() {
	// https://stackoverflow.com/questions/22535775/the-last-element-of-a-slice
	return stack[:len(sl)-1]
}

func save_state(d deck.Deck) {
	stack = 
}

func redo(stack []deck.Deck) {
	return sl[:len(sl)-1]
}
