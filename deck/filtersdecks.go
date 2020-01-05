package deck

type DeckControls []*DeckControl

// Writes multiple deck controls
func (decks DeckControls) Write() error {
	for _, d := range decks {
		err := d.Write()
		if err != nil {
			return err
		}
	}

	return nil
}

func (decks DeckControls) FilterGroups(param []string) (out_deck Deck) {
	for _, d := range decks {
		out_deck = append(out_deck, d.Deck.FilterGroups(param)...)
	}
	return
}

func (decks DeckControls) FilterGroupsAdd(param []string) (out_deck Deck) {
	for _, d := range decks {
		out_deck = append(out_deck, d.Deck.FilterGroupsAdd(param)...)
	}
	return
}

func (decks DeckControls) FilterGroup(param string) (out_deck Deck) {
	for _, d := range decks {
		out_deck = append(out_deck, d.Deck.FilterGroup(param)...)
	}
	return
}

func (decks DeckControls) FilterReview() (out_deck Deck) {
	for _, d := range decks {
		out_deck = append(out_deck, d.Deck.FilterReview()...)
	}
	return
}

func (decks DeckControls) FilterDone() (out_deck Deck) {
	for _, d := range decks {
		out_deck = append(out_deck, d.Deck.FilterDone()...)
	}
	return
}

func (decks DeckControls) FilterMemorize() (out_deck Deck) {
	for _, d := range decks {
		out_deck = append(out_deck, d.Deck.FilterMemorize()...)
	}
	return
}

func OpenDecks(filenames []string) (DeckControls, error) {
	var dcks DeckControls

	for _, filename := range filenames {
		d, err := OpenNewFormat(filename)
		if err != nil {
			return nil, err
		} else {
			dcks = append(dcks, d)
		}
	}

	return dcks, nil
}
