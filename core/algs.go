package core

import (
	"fmt"
	"time"
)

// Implementing the SM2 algorithm.
func sm2Exec(s Meta, input Know) *Meta {
	// Input defaults.
	var rank float32 = 2.5
	if len(s.Params) > 0 {
		rank = floatOrDefault(s.Params[0], rank)
	}

	// Sm2 Logic
	var q float32
	s.Next = time.Now()

	if input == NO {
		q = 1
	} else if input == IDK {
		q = 3
	} else if input == YES {
		q = 5
	}

	if s.Streak > 0 {
		rank += -.8 + .28*q - .02*q*q
	}

	if rank < 1.3 {
		rank = 1.3
	}

	if input == YES {
		nextDay := float32(1.0)

		if s.Streak < 0 {
			s.Streak = 0
		}

		if s.Streak >= 1 {
			nextDay += 5
		}

		if s.Streak >= 2 {
			for i := 2; i <= s.Streak; i++ {
				nextDay *= rank
			}
		}

		s.Streak++
		s.Next = time.Now().AddDate(0, 0, int(nextDay))
	} else {
		if s.Streak > 0 {
			s.Streak = -1
		} else if s.Streak < 0 {
			s.Streak -= 1
		}
	}

	s.Params = []string{fmt.Sprintf("%.2f", rank)}
	return &s
}
