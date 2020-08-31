package core

import "fmt"

type MetaAlg struct {
   MetaBase
   Name   string
}

func NewMetaAlgFromStrings(strs []string) *MetaAlg {
   return &MetaAlg {
      MetaBase: *NewMetaBase(strs),
      Name: getParam(strs, 6),
   }
}

func NewDefaultMetaAlg(hash string, name string) *MetaAlg {
   return &MetaAlg {
      MetaBase: *NewMetaBase([]string{hash}),
      Name: name,
   }
}

func NewMetaAlg(ai *AlgInfo, mh *MetaHist) *MetaAlg {
   // Yes/No count
   yesCount := mh.YesCount
   noCount := mh.NoCount
   if mh.Target {
      yesCount++
   } else {
      noCount++
   }

   // Streak Logic
   streak := mh.Streak
   switch mh.GetAnswerCategory() {
      case YesWasYes: streak++
      case NoWasNo:   streak--
      default: streak=0
   }

   return &MetaAlg{
      MetaBase{
         mh.Hash,
         ai.Next,
         mh.Next,
         yesCount,
         noCount,
         streak,
      }, ai.Name,
   }
}

func (m *MetaAlg) Exec(input bool) (*MetaAlg, error) {
   mh := NewMetaHistFromMetaAlg(m, input)

   // Save the current time for logging & not saving the current time multiple times.
   var ai AlgInfo
   if algFunc, exists := Algs[m.Name]; exists {
      ai = algFunc(*mh)
   } else {
      return nil, fmt.Errorf("Algorithm doesn't exist.")
   }

   return NewMetaAlg(&ai, mh), nil
}

func (m *MetaAlg) IsZero() bool {
   return m.MetaBase.IsZero() && m.Name == ""
}

func (m *MetaAlg) String() (s string) {
   if !m.IsZero() {
      s = fmt.Sprintf("%s %s", m.MetaBase.String(), m.Name)
   }

   return
}
