package meta

import (
	"math"
	"math/rand"
	"time"
)

// Modified SM2 Algorithm
// Returns the duration in nanoseconds for when to review the card next.
func sm2Exec(r Outcome) float64 {
	period := 0.0
	exponent := 0
	rank := math.Max(1.3, 2.5 + .11*float64(r.YesCount()) - .20*float64(r.NoCount()) + .08*float64(r.Streak()))
	// Next Day Logic
	if r.Target() && r.Streak() < 0 {
		period = float64(time.Minute * 2)
	} else if r.Target() {
		exponent = r.Streak()
		if r.Streak() == 0 {
			period = float64(time.Minute * 4)
		} else if r.Streak() == 1 {
			period = float64(time.Hour * 4)
		} else if r.Streak() > 1 {
			period = float64(time.Hour * 8)
		}
	} else {
		period = float64(time.Minute)
	}

	// Multiply by the rank x number of times.
	for i := 0; i <= exponent; i++ {
		period *= rank
	}

	// Add some noise, so everything doesn't get reviewed at the same time.
	return period * (1 + .1*rand.Float64())
}
