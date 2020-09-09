package core

import "time"
import "fmt"

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

func NewMetaHistFromStrings(strs ...string) *MetaHist {
   return &MetaHist{
      metaBase: *newMetaBaseFromStrings(strs...),
      Target: getParam(strs, 6) == "1",
   }
}

func NewMetaHistFromMetaAlg(ma *MetaAlg, target bool) *MetaHist {
   mh := &MetaHist{
      metaBase: ma.metaBase,
      Target: target,
   }

   mh.next = time.Now()
   if mh.curr.IsZero() {
      mh.curr = mh.next
   }

   return mh
}

func (mh *MetaHist) AnswerCategory() AnswerCategory {
   if mh.Target {
      if mh.streak < 0 {
         return YesWasNo
      } else {
         return YesWasYes
      }
   } else {
      if mh.streak > 0 {
         return NoWasYes
      } else {
         return NoWasNo
      }
   }
}

func (mh *MetaHist) NewStreak() int {
   // Streak Logic
   streak := mh.streak
   switch mh.AnswerCategory() {
      case YesWasYes: streak++
      case NoWasNo:   streak--
      default: streak=0
   }
   return streak
}

func (mh *MetaHist) newCount(expecting bool, count int) int {
   if expecting == mh.Target { count++ }
   return count
}

func (mh *MetaHist) NewYesCount() int { return mh.newCount(true, mh.yesCount) }
func (mh *MetaHist) NewNoCount()  int { return mh.newCount(false, mh.noCount) }
func (mh *MetaHist) TargetStr() string { if mh.Target { return "1" } else { return "0" } }

func (mh *MetaHist) String() string {
   return fmt.Sprintf("%s %s", mh.metaBase.String(), mh.TargetStr())
}
