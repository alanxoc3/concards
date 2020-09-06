package core

import "time"

type AnswerCategory uint8

const (
   YesWasYes AnswerCategory = 1 << iota
   YesWasNo
   NoWasYes
   NoWasNo
)

type MetaHist struct {
   metaBase
   Target bool
}

func NewMetaHistFromMetaAlg(ma *MetaAlg, target bool) *MetaHist {
   mh := &MetaHist{
      metaBase: ma.metaBase,
      Target: target,
   }

   mh.Next = time.Now()
   return mh
}

func (m *MetaHist) GetAnswerCategory() AnswerCategory {
   if m.Target {
      if m.Streak < 0 {
         return YesWasNo
      } else {
         return YesWasYes
      }
   } else {
      if m.Streak > 0 {
         return NoWasYes
      } else {
         return NoWasNo
      }
   }
}
