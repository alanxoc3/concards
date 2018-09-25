package algs

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/alanxoc3/concards/constring"
)

type SpaceAlg struct {
	Next   time.Time
	Streak int
	Rank   float32
}

func (s *SpaceAlg) ToString() string {
	next := constring.DateToString(s.Next)

	str := fmt.Sprintf("%s, %d, %.2f", next, s.Streak, s.Rank)
	return str
}

func New(str string) (*SpaceAlg, error) {
	// Split the string into bins, then put the bins in the correct spots.
	space := SpaceAlg{Next: time.Now(), Streak: 0, Rank: 2.5}

	bins := strings.Split(str, ",")
	binsLen := len(bins)

	// trim the bins
	for i := 0; i < binsLen; i++ {
		bins[i] = constring.Trim(bins[i])
	}

	if binsLen > 0 && !constring.IsEmpty(bins[0]) { // NEXT
		x, err := constring.StrToDate(bins[0])
		if err != nil {
			return nil, err //errors.New("Problem parsing meta data.")
		}
		space.Next = x
	}

	if binsLen > 1 && !constring.IsEmpty(bins[1]) { // STREAK
		x, err := strconv.Atoi(bins[1])
		if err != nil {
			return nil, errors.New("Problem parsing meta data.")
		}
		space.Streak = x
	}

	if binsLen > 2 && !constring.IsEmpty(bins[2]) { // STREAK
		x, err := strconv.ParseFloat(bins[2], 32)
		if err != nil {
			return nil, errors.New("Problem parsing meta data.")
		}
		space.Rank = float32(x)
	}

	return &space, nil
}

/*
	Spaced algorithm has these values.
	Next date - Next time algorithm should be used
	Streak    - how many times you've gotten it right.
	Rank    - a number specific to the algorithm.
*/
