package card

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"crypto/sha256"

	"github.com/alanxoc3/concards/internal"
)

// A card is a list of facts. Usually, but not limited to, Q&A format.
type Card struct {
	file  string
	facts []string
}

type CardMap map[internal.Hash]*Card

func splitBySide(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Check for "::".
	if len(data) >= 2 {
		if data[0] == byte(':') && data[1] == byte(':') {
			return 2, data[:2], nil
		}
	}

	// Check for "|".
	if len(data) >= 1 {
		if data[0] == byte('|') {
			return 1, data[:1], nil
		}
	}

	// Parse until next token
	isBackslash := false
	isColon := false
	for width, i := 0, 0; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])

		if !isBackslash {
			if r == '|' {
				return i, data[:i], nil
			} else if isColon && r == ':' {
				return i - 1, data[:i-1], nil
			} else if r == ':' {
				isColon = true
			} else {
				isColon = false
			}
		} else {
			isColon = false
		}

		isBackslash = r == '\\' && !isBackslash
	}

	// Return the non empty remainder.
	if atEOF && len(data) > 0 {
		return len(data), data[:], nil
	}

	// Request more data.
	return 0, nil, nil
}

func splitByWordBackslash(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(r) {
			break
		}
	}

	// Parse until next token
	isBackslash := false
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])

		if !isBackslash {
			if unicode.IsSpace(r) {
				return i, data[start:i], nil
			}
		}

		isBackslash = r == '\\' && !isBackslash
	}

	// Return the non empty remainder.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}

	// Request more data.
	return start, nil, nil
}


func splitByToken(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Check for "::" or "{}".
	if len(data) >= 2 {
		if data[0] == byte('{') && data[1] == byte('}') {
			return 2, data[:2], nil
		}
	}

	// Check for "|" or "{" or "}" or ":" or "#".
	if len(data) >= 1 {
		if data[0] == byte('#') || data[0] == byte(':') || data[0] == byte('{') || data[0] == byte('}') {
			return 1, data[:1], nil
		}
	}

	// Parse until next token
	isBackslash := false
	for width, i := 0, 0; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])

		if !isBackslash {
			if r == '{' || r == '}' || r == '#' || r == ':' {
				return i, data[0:i], nil
			}
		}

		isBackslash = r == '\\' && !isBackslash
	}

	// Return the non empty remainder.
	if atEOF && len(data) > 0 {
		return len(data), data[:], nil
	}

	// Request more data.
	return 0, nil, nil
}

// Unbackslashes things that don't need to be backslashed.
func normalizeBackslash(side string) string {
	s := []rune{}
   isBackslash := false

   for _, r := range side {
      if isBackslash {
         if r == '\\' || r == '#' || r == '{' || r == '}' || r == ':' || r == '|' || unicode.IsSpace(r) {
            s = append(s, '\\', r)
         } else {
            s = append(s, r)
         }
      } else if r != '\\' {
         s = append(s, r)
      }

		isBackslash = r == '\\' && !isBackslash
	}

	return string(s)
}

// Backslashes any invalid cloze or colon.
func normalizeCloze(side string) string {
	s := ""
	clozeString := ""
	clozeDepth := 0

	scanner := bufio.NewScanner(strings.NewReader(side))
	scanner.Split(splitByToken)
	for scanner.Scan() {
		t := scanner.Text()

		if clozeDepth > 0 {
			if t == "{" || t == "}" || t == ":" || t == "#" {
				clozeString += t
				s += "\\" + t
			} else {
				clozeString += t
				s += t
			}

			if t == "{" {
				clozeDepth++
			} else if t == "}" {
				clozeDepth--
			}

			if clozeDepth == 0 {
				s = clozeString
			}
		} else {
			if t == "{" {
				clozeDepth++
				clozeString = s + "{"
				s += "\\{"
			} else if t == "}" || t == ":" {
				s += "\\" + t
			} else {
				s += t
			}
		}
	}

	return s
}

// Backslashes any hashes not associated with a cloze.
func normalizeHash(side string) string {
	s := ""
	hashCount := 0

	scanner := bufio.NewScanner(strings.NewReader(side))
	scanner.Split(splitByToken)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "#" {
			hashCount++
		} else {
			if hashCount > 0 && t == "{" {
				s += strings.Repeat("#", hashCount)
			} else if hashCount > 0 {
				s += strings.Repeat("\\#", hashCount)
			}
			s += t
			hashCount = 0
		}
	}

	if hashCount > 0 {
		s += strings.Repeat("\\#", hashCount)
	}

	return s
}

// Turns colons into curly braces.
func normalizeColon(side string) string {
	s := ""
	hashCount := 0
	groupStack := []int{}

	scanner := bufio.NewScanner(strings.NewReader(side))
	scanner.Split(splitByToken)
	for scanner.Scan() {
		t := scanner.Text()

		if t == ":" {
			if groupStack[len(groupStack)-1] == 0 {
				s += "}{"
			}
		} else {
			s += t
		}

		if t == "#" {
			hashCount++
		} else if t == "{" {
			groupStack = append(groupStack, hashCount)
			hashCount = 0
		} else if t == "}" {
			groupStack = groupStack[:len(groupStack)-1]
		}
	}

	return s
}

type clozeNode struct {
	loc   int
	group int
	text  string
	nodes []*clozeNode
}

func trimSpacesWithBackslash(side string) string {
	words := []string{}

	scanner := bufio.NewScanner(strings.NewReader(side))
	scanner.Split(splitByWordBackslash)
	for scanner.Scan() {
      words = append(words, scanner.Text())
	}

	return strings.Join(words, " ")
}

// Turns colons into curly braces.
func calcClozeNode(scanner *bufio.Scanner) *clozeNode {
	nodeText := ""
	nodes := []*clozeNode{}
	hashCount := 0

	for scanner.Scan() {
		t := scanner.Text()
		if t == "{" {
			node := calcClozeNode(scanner)
			node.loc = len(nodeText)
			node.group = hashCount
			nodes = append(nodes, node)

			hashCount = 0
		} else if t == "#" {
			hashCount++
		} else if t == "}" {
			break
		} else {
			hashCount = 0
			nodeText += t
		}
	}

	return &clozeNode{
		loc:   0,
		group: 0,
		text:  nodeText,
		nodes: nodes,
	}
}

// Returns a list of cards, or an empty list if there is an error.
func NewCards(file string, cardStr string) ([]*Card, error) {
	if file == "" {
		return []*Card{}, fmt.Errorf("File not provided.")
	}

	cards := []*Card{}
	sides := []string{}
	revSides := [][]string{}

	// Step 1: Scan through curly braces.

	// Step 1: Scan through string by card words and special tokens.
	scanner := bufio.NewScanner(strings.NewReader(cardStr))
	scanner.Split(splitBySide)
	for scanner.Scan() {
		side := scanner.Text()
		if side == "::" {
			if len(sides) > 0 {
				revSides = append(revSides, sides)
				sides = []string{}
			}
		} else if side != "|" {
			side = normalizeBackslash(side)
			side = normalizeCloze(side)
			side = normalizeHash(side)
			side = normalizeColon(side)
			side = trimSpacesWithBackslash(side)

			tokenScanner := bufio.NewScanner(strings.NewReader(side))
			tokenScanner.Split(splitByToken)
			n := calcClozeNode(tokenScanner)

         if len(n.text) > 0 {
            sides = append(sides, n.text)
         }
		}
	}

	// Step 2: Put any remaining sides to the reverse card structure.
	if len(sides) > 0 {
		revSides = append(revSides, sides)
	}


	// Step 3: Add all the cards/reverse cards.
	if len(revSides) > 1 {
		for ri, rs := range revSides {
			for _, s := range rs {
				facts := []string{}
				facts = append(facts, s)
				for rri, rrs := range revSides {
					if rri != ri {
						for _, ss := range rrs {
							facts = append(facts, ss)
						}
					}
				}

				cards = append(cards, &Card{file, facts})
			}
		}
	} else if len(revSides) == 1 {
		cards = append(cards, &Card{file, revSides[0]})
	}

	return cards, nil
}

func (c *Card) HasAnswer() bool { return len(c.facts) > 1 }
func (c *Card) String() string  { return strings.Join(c.facts, " | ") }
func (c *Card) File() string    { return c.file }
func (c *Card) Len() int        { return len(c.facts) }

func (c *Card) Hash() (dest internal.Hash) {
	hash := sha256.Sum256([]byte(c.String()))
	copy(dest[:], hash[:])
	return dest
}

func (c *Card) GetFactEsc(i int) string {
	factStr := c.getFactRaw(i)

	scanner := bufio.NewScanner(strings.NewReader(factStr))
	scanner.Split(bufio.ScanRunes)
	prev := ""
	str := ""
	for scanner.Scan() {
		t := scanner.Text()

		if prev == "\\" {
			str += t
		} else if t != "\\" {
			str += t
		}
		prev = t
	}
	return str
}

// -------------------- Private stuff below --------------------

func (c *Card) getFactRaw(i int) string {
	if len(c.facts) > i && 0 <= i {
		return c.facts[i]
	}
	return ""
}
