package card

// Card represents a single flash card. Contains all
// information pertaining to a card.
type Card struct {
	Question string
	Answer   string
	Groups   []string
	Metadata interface{}
}
