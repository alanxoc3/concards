package deck

import (
	"time"

	"github.com/alanxoc3/concards-go/constring"
)

func (cards Deck) FilterFile(param string) Deck {
	var list Deck

	for _, c := range cards {
		if c.File == param {
			list = append(list, c)
		}
	}

	return list
}

func (cards Deck) FilterGroups(param []string) Deck {
	var list Deck

	for _, c := range cards {
		if constring.ListHasOtherList(c.Groups, param) {
			list = append(list, c)
		}
	}

	return list
}

func (cards Deck) FilterGroup(param string) Deck {
	var groups []string
	groups = append(groups, param)
	return cards.FilterGroups(groups)
}

func (cards Deck) FilterReview() Deck {
	var list Deck

	for _, c := range cards {
		if c.Metadata.Next.Before(time.Now()) && c.Metadata.Streak != 0 {
			list = append(list, c)
		}
	}

	return list
}

func (cards Deck) FilterDone() Deck {
	var list Deck

	for _, c := range cards {
		if c.Metadata.Next.After(time.Now()) && c.Metadata.Streak != 0 {
			list = append(list, c)
		}
	}

	return list
}

func (cards Deck) FilterMemorize() Deck {
	var list Deck

	for _, c := range cards {
		if c.Metadata.Streak == 0 {
			list = append(list, c)
		}
	}

	return list
}
