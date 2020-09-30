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
	rank := math.Max(1.3, 2.5+.11*float64(r.YesCount())-.29*float64(r.NoCount())+.06*float64(r.Streak()))

	// Next Day Logic
   if r.Target() {
      if r.Streak() < 0 {
         period = float64(time.Minute * 5)
      } else if r.Streak() == 0 {
         period += float64(time.Hour * 24)
      } else if r.Streak() > 0 {
         period += float64(time.Hour * 24 * 6)
      }

      if r.Streak() >= 2 {
         for i := 2; i <= r.Streak(); i++ {
            period *= rank
         }
      }
   } else {
		period = float64(time.Minute * 1)
   }

	// Add some noise, so everything doesn't get reviewed at the same time.
	return period * (1 + .1*rand.Float64())
}
