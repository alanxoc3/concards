package meta

import (
	"math"
	"math/rand"
	"time"
)

// Constants
const MaxNextDuration float64 = float64(time.Hour * 24 * 365 * 100) // 100 years.
const MaxYesNoStreak = 1 << 30                                      // About a billion.
type AnswerClassification uint8
type AlgFunc func(Result) float64

const (
	YesWasYes AnswerClassification = 1 << iota
	YesWasNo
	NoWasYes
	NoWasNo
)

// TODO: Remove this with GH-45.
var algs = map[string]AlgFunc{
	"sm2": sm2Exec,
}

// Modified SM2 Algorithm
// Returns the duration in nanoseconds for when to review the card next.
func sm2Exec(r Result) float64 {
	ac := r.AnswerClassification()
	period := 0.0
	rank := math.Max(1.3, 2.5+.11*float64(r.YesCount())-.29*float64(r.NoCount())+.06*float64(r.Streak()))

	// Next Day Logic
	if ac == YesWasYes {
		if r.Streak() < 0 {
			panic("Logic error with concards! Please make an issue on github.")
		} else if r.Streak() == 0 {
			period += float64(time.Hour * 24)
		} else {
			period += float64(time.Hour * 24 * 6)
		}

		if r.Streak() >= 2 {
			for i := 2; i <= r.Streak(); i++ {
				period *= rank
			}
		}
	} else if ac == YesWasNo {
		period = float64(time.Minute * 5)
	} else {
		period = float64(time.Minute * 1)
	}

	// Add some noise, so everything doesn't get reviewed at the same time.
	return period * (1 + .1*rand.Float64())
}
