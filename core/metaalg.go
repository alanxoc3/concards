package core

import "fmt"

type MetaAlg struct {
   metaBase
   Name   string
}

func NewMetaAlgFromStrings(strs ...string) *MetaAlg {
   return &MetaAlg {
      metaBase: *NewMetaBaseFromStrings(strs...),
      Name: getParam(strs, 6),
   }
}

func NewDefaultMetaAlg(hash string, name string) *MetaAlg {
   return &MetaAlg {
      metaBase: *NewMetaBaseFromStrings([]string{hash}...),
      Name: name,
   }
}

func newMetaAlg(ai *AlgInfo, mh *MetaHist) *MetaAlg {
   return &MetaAlg{
      *newMetaBase(
         mh.hash,
         ai.Next,
         mh.next,
         mh.NewYesCount(),
         mh.NewNoCount(),
         mh.NewStreak(),
      ), ai.Name,
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

   return newMetaAlg(&ai, mh), nil
}

func (m *MetaAlg) String() string {
   return fmt.Sprintf("%s %s", m.metaBase.String(), m.Name)
}
