package algs

import "time"

type Know uint16

const (
	NO  Know = iota
	IDK      = iota
	YES      = iota
)

// Implementing the SM2 algorithm.
func (s *SpaceAlg) Execute(input Know) {
	// If today is after the next day, then we can review the card.
	if time.Now().Before(s.Next) && input == YES {
		return
	}

	// Validate input.
	if input != NO && input != IDK && input != YES {
		return
	}

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
		s.Rank += -.8 + .28*q - .02*q*q
	}

	if s.Rank < 1.3 {
		s.Rank = 1.3
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
				nextDay *= s.Rank
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
}
