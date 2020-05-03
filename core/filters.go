package core

// Basically truncates a deck.
func (d Deck) FilterNumber(param int) Deck {
	if param > 0 && d.Len() > param {
      d.refs = d.refs[0:param]
	}
   return d
}
