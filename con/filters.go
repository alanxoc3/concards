package con

// Trims a deck to only have this many d in it.
func (d Deck) FilterNumber(param int) Deck {
	if param > 0 && d.Len() > param {
      d.refs = d.refs[0:param]
	}
   return d
}

// Trims a deck to only contain these groups in it.
// func (d Deck) FilterGroups(param []string) Deck {
	// var list Deck
// 
	// // for _, c := range d {
		// // if constring.ListHasOtherList(c.Groups, param) {
			// // list = append(list, c)
		// // }
	// // }
// 
	// return list
// }

// Trims a deck to only contain these groups in it.
// func (d Deck) FilterGroup(param string) Deck {
	// var groups []string
	// groups = append(groups, param)
	// return d.FilterGroups(groups)
// }

// idk
// func (d Deck) FilterGroupsAdd(param []string) Deck {
	// var list Deck
// 
	// for _, p := range param {
		// list = append(list, d.FilterGroup(p)...)
	// }
// 
	// return list
// }

// func (d Deck) FilterReview() Deck {
// 	var list Deck
// 
// 	for _, c := range d {
// 		if c.Metadata.Next.Before(time.Now()) && c.Metadata.Streak != 0 {
// 			list = append(list, c)
// 		}
// 	}
// 
// 	return list
// }
// 
// func (d Deck) FilterDone() Deck {
// 	var list Deck
// 
// 	for _, c := range d {
// 		if c.Metadata.Next.After(time.Now()) && c.Metadata.Streak != 0 {
// 			list = append(list, c)
// 		}
// 	}
// 
// 	return list
// }
// 
// func (d Deck) FilterMemorize() Deck {
// 	var list Deck
// 
// 	for _, c := range d {
// 		if c.Metadata.Streak == 0 {
// 			list = append(list, c)
// 		}
// 	}
// 
// 	return list
// }
