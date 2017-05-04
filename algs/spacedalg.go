package algs

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/alanxoc3/concards-go/constring"
)

type SpaceAlg struct {
	name   string
	next   time.Time
	first  time.Time
	streak int
	high   int
	low    int
	extra  []string
}

func (s *SpaceAlg) ToString() string {
	next := constring.DateToString(s.next)
	first := constring.DateToString(s.first)
	extra := constring.FormatList(s.extra)

	str := fmt.Sprintf("%s, %s, %s, %d, %d, %d%s", s.name, next, first, s.streak, s.high, s.low, extra)
	return str
}

func New(str string) (*SpaceAlg, error) {
	// Split the string into bins, then put the bins in the correct spots.
	space := SpaceAlg{}
	bins := strings.Split(str, ",")
	binsLen := len(bins)

	// trim the bins
	for i := 0; i < binsLen; i++ {
		bins[i] = constring.Trim(bins[i])
	}

	space.name = bins[0] // NAME
	bins[0] = strings.ToUpper(bins[0])

	if bins[0] == "" {
		space.name = "SM2"
	} else if IsValidAlgName(bins[0]) {
		space.name = bins[0]
	} else {
		return nil, errors.New("Invalid Algorithm Name")
	}

	if binsLen > 1 && !constring.IsEmpty(bins[1]) { // NEXT
		x, err := constring.StrToDate(bins[1])
		if err != nil {
			return nil, err
		}
		space.next = x
	}

	if binsLen > 2 && !constring.IsEmpty(bins[2]) { // FIRST
		x, err := constring.StrToDate(bins[2])
		if err != nil {
			return nil, err
		}
		space.first = x
	}

	if binsLen > 3 && !constring.IsEmpty(bins[3]) { // STREAK
		x, err := strconv.Atoi(bins[3])
		if err != nil {
			return nil, err
		}
		space.streak = x
	}

	if binsLen > 4 && !constring.IsEmpty(bins[4]) { // HIGH
		x, err := strconv.Atoi(bins[4])
		if err != nil {
			return nil, err
		}
		space.high = x
	}

	if binsLen > 5 && !constring.IsEmpty(bins[5]) { // LOW
		x, err := strconv.Atoi(bins[5])
		if err != nil {
			return nil, err
		}
		space.low = x
	}

	if binsLen > 6 { // EXTRA
		space.extra = bins[6:]
	}

	return &space, nil
}

/*
Spaced Algorithms explained:
All Spaced Algorithms have certain properties that are the same between
them. These properties are:
	name   = the name of the algorithm.
	next   = the next date the card needs to be reviewed.
	first  = initial date the card was passed
	streak = the current number of times you have gotten the card correct.
	high   = high score streak
	low    = low score streak
	extra  = extra data specific to the algorithm.

	Each algorithm must worry about all these variables in order to work correctly.

*/

func IsValidAlgName(name string) bool {
	return name == "SM2" || name == "PRATT" || name == "DAILY" || name == "SHELL"
}
