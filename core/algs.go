package core

import (
   "fmt"
   "time"
   "math"
   "math/rand"
	"strconv"
)

type AlgInfo struct {
   Next   time.Time
   Name   string
   Params []string
}

func floatOrDefault(str string, def float64) float64 {
	if x, err := strconv.ParseFloat(str, 64); err != nil {
		return def
	} else {
		return float64(x)
	}
}

// SM2 Algorithm
// Returns meta & location to put card in deck.
func sm2Exec(mh MetaHist, ma MetaAlg) *AlgInfo {
   const maxPeriod float64 = float64(time.Hour*24*365*100)
   const randPercentage float64 = .1

   ac := mh.GetAnswerCategory()
   period := 0.0
   rank := 2.5

   // Rank Logic
   if len(ma.Params) > 0 {
      rank = floatOrDefault(ma.Params[0], rank)
   }

   switch ac {
      case YesWasYes: rank += .10
      case YesWasNo:  rank += .03
      case NoWasYes:  rank -= .32
      case NoWasNo:   rank -= .05
   }

   rank = math.Max(1.3, rank)

   // Next Day Logic
   if ac == YesWasYes {
      if ma.Streak < 0 {
         panic("Logic error with concards! Please make an issue on github.")
      } else if ma.Streak == 0 {
         period += float64(time.Hour*24)
      } else {
         period += float64(time.Hour*24*6)
      }

      if ma.Streak >= 2 {
         for i := 2; i <= ma.Streak; i++ {
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
   return &AlgInfo{
      Next: mh.Next.Add(time.Duration(period)),
      Name: "sm2",
      Params: []string{fmt.Sprintf("%.2f", rank)}
   }
}
