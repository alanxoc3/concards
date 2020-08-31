package core

import (
   "time"
   "math"
   "math/rand"
)

type AlgFunc func(MetaHist) AlgInfo

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
// Returns meta & location to put card in deck.
func sm2Exec(mh MetaHist) AlgInfo {
   const maxPeriod float64 = float64(time.Hour*24*365*100)
   const randPercentage float64 = .1

   ac := mh.GetAnswerCategory()
   period := 0.0
   rank := math.Max(1.3, 2.5 + .1*float64(mh.YesCount) - .3*float64(mh.NoCount) + .05*float64(mh.Streak))

   // Next Day Logic
   if ac == YesWasYes {
      if mh.Streak < 0 {
         panic("Logic error with concards! Please make an issue on github.")
      } else if mh.Streak == 0 {
         period += float64(time.Hour*24)
      } else {
         period += float64(time.Hour*24*6)
      }

      if mh.Streak >= 2 {
         for i := 2; i <= mh.Streak; i++ {
            period *= rank
         }
      }
   } else if ac == YesWasNo {
      period = float64(time.Minute*5)
   } else {
      period = float64(time.Minute*1)
   }

   period = math.Min(period * (1 + rand.Float64()*randPercentage), maxPeriod)

   // The "Next" on meta history should represent "time.Now".
   return AlgInfo{
      Next: mh.Next.Add(time.Duration(period)),
      Name: "sm2",
   }
}
