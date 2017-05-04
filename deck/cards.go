package deck

import (
	"sort"

	"github.com/alanxoc3/concards-go/card"
	"github.com/alanxoc3/concards-go/constring"
)

type Cards []*card.Card

func (cards Cards) Len() int {
	return len(cards)
}

func (cards Cards) Less(i, j int) bool {
	return cards[i].Id < cards[j].Id
}

func (cards Cards) Swap(i, j int) {
	cards[i], cards[j] = cards[j], cards[i]
}

func (cards Cards) Sort() {
	sort.Sort(cards)
}

// includes the group markers
func (cards Cards) ToString() string {
	// do groups stuff
	str := ""
	var curGroups []string

	for _, c := range cards {
		if !constring.StringListsIdentical(curGroups, c.Groups) {
			curGroups = c.Groups
			str += constring.ListToString(curGroups) + "\n"
		}

		str += c.ToString() + "\n\n"
	}

	return str
}

func (cards Cards) FilterFile(param string) Cards {
	var list Cards

	for _, c := range cards {
		if c.File == param {
			list = append(list, c)
		}
	}

	return list
}

func (cards Cards) FilterGroups(param []string) Cards {
	var list Cards

	for _, c := range cards {
		if constring.ListHasOtherList(c.Groups, param) {
			list = append(list, c)
		}
	}

	return list
}

func (cards Cards) FilterGroup(param string) Cards {
	var groups []string
	groups = append(groups, param)
	return cards.FilterGroups(groups)
}
