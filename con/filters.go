package con

// Trims a deck to only have this many d in it.
func (d Deck) FilterNumber(param int) Deck {
	if param > 0 && d.Len() > param {
      d.refs = d.refs[0:param]
	}
   return d
}
