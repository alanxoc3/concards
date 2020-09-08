package core

import "fmt"
import "time"
import "math"

// A given interval may not exceed 100 years.
// This is partially limited by Go's time.Duration.
const MaxInterval float64 = float64(time.Hour*24*365*100)

type MetaAlg struct {
   metaBase
   name   string
}

func NewMetaAlgFromStrings(strs ...string) *MetaAlg {
   return &MetaAlg {
      metaBase: *newMetaBaseFromStrings(strs...),
      name: getParam(strs, 6),
   }
}

func NewDefaultMetaAlg(hash string, name string) *MetaAlg {
   return &MetaAlg {
      metaBase: *newMetaBaseFromStrings([]string{hash}...),
      name: name,
   }
}

func (ma *MetaAlg) Exec(input bool) (*MetaAlg, error) {
   // Note that mh.Next() has the current time.
   mh := NewMetaHistFromMetaAlg(ma, input)

   var next time.Time
   if algFunc, exists := Algs[ma.name]; exists {
      next = mh.Next().Add(time.Duration(math.Min(algFunc(*mh), MaxInterval)))
   } else {
      return nil, fmt.Errorf("Algorithm doesn't exist.")
   }

   return &MetaAlg{
      *newMetaBase(
         mh.Hash(),
         next,
         mh.Next(),
         mh.NewYesCount(),
         mh.NewNoCount(),
         mh.NewStreak(),
      ), ma.Name(),
   }, nil
}

func (m *MetaAlg) Name() string {
   return m.name
}

func (m *MetaAlg) String() string {
   return fmt.Sprintf("%s %s", m.metaBase.String(), m.name)
}
