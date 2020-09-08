package core

import (
   "time"
   "math"
   "math/rand"
)

type AlgFunc func(MetaHist) float64

var Algs = map[string]AlgFunc{
   "sm2": sm2Exec,
}

type AlgInfo struct {
   Next   time.Time
   Name   string
}

// If streak = 0, curr 0. prev + or -.
// If streak > 0, curr +. prev 0 or +.
// If streak < 0, curr -. prev 0 or -.

// SM2 Algorithm
// Returns the duration in nanoseconds for when to review the card next.
func sm2Exec(mh MetaHist) float64 {
   ac := mh.AnswerCategory()
   period := 0.0
   rank := math.Max(1.3, 2.5 + .1*float64(mh.YesCount()) - .3*float64(mh.NoCount()) + .05*float64(mh.Streak()))

   // Next Day Logic
   if ac == YesWasYes {
      if mh.Streak() < 0 {
         panic("Logic error with concards! Please make an issue on github.")
      } else if mh.Streak() == 0 {
         period += float64(time.Hour*24)
      } else {
         period += float64(time.Hour*24*6)
      }

      if mh.Streak() >= 2 {
         for i := 2; i <= mh.Streak(); i++ {
            period *= rank
         }
      }
   } else if ac == YesWasNo {
      period = float64(time.Minute*5)
   } else {
      period = float64(time.Minute*1)
   }

   // Add some noise, so everything doesn't get reviewed at the same time.
   return period * (1 + .1*rand.Float64())
}
