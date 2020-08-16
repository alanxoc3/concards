package core

import (
   "fmt"
   "time"
   "math"
   "math/rand"
)

// SM2 Algorithm
// Returns meta & location to put card in deck.
func sm2Exec(m MetaHist) *MetaAlg {
   ac := m.GetAnswerCategory(input)

   // Rank Logic
   var rank float32 = 2.5
   if len(m.Params) > 0 {
      rank = floatOrDefault(m.Params[0], rank)
   }

   switch ac {
      case YesWasYes: rank += .10
      case YesWasNo:  rank += .03
      case NoWasYes:  rank -= .32
      case NoWasNo:   rank -= .05
   }

   rank = math.Max(1.3, rank)
   m.Params = []string{fmt.Sprintf("%.2f", rank)}

   // Next Day Logic
   if ac == YesWasYes {
      nextDay := float32(1.0)
      if m.Streak < 0 { panic("Logic error with concards! Please make an issue on github.") }

      if m.Streak > 0 {
         nextDay += 5
      }

      if m.Streak >= 2 {
         for i := 2; i <= m.Streak; i++ {
            nextDay *= rank
         }
      }

      // 3 extra days for randomness.
      m.Next = m.Next.Add(time.Day*nextDay + time.Second * rand.Intn(86400*3))
   } else if ac == YesWasNo {
      m.Next = m.Next.Add(time.Minute*5 + time.Second*rand.Intn(120))
   } else {
      m.Next = m.Next.Add(time.Minute + time.Second*rand.Intn(60))
   }

   return m
}
