package core

import "time"
import "fmt"
import "strings"

type MetaAlg struct {
   *MetaBase
   Name   string
   Params []string
}

func NewMetaAlgFromStrings(next string, curr string, streak string, name string, params []string) *MetaAlg {
   return &MetaAlg {
      MetaBase: NewMetaBaseFromStrings(next, curr, streak),
      Name: name,
      Params: params,
   }
}

func NewDefaultMetaAlg(name string) *MetaAlg {
   return &MetaAlg {
      MetaBase: &MetaBase{time.Now(), time.Now(), 0},
      Name: name,
      Params: []string{},
   }
}

func (m *MetaAlg) Exec(input bool) (*MetaAlg, error) {
   ma := &MetaAlg{}
   *ma = *m
   mh := NewMetaHistFromMetaAlg(m)

   // Save the current time for logging & not saving the current time multiple times.
   switch ma.Name {
      case "sm2": ma = sm2Exec(*mh, *m)
      default: return m, fmt.Errorf("Algorithm doesn't exist.")
   }

   // Streak Logic
   switch ma.GetAnswerCategory() {
      case YesWasYes: ma.Streak++
      case NoWasNo:   ma.Streak--
      default: ma.Streak=0
   }

   return ma, nil
}

func (m *MetaAlg) IsZero() bool {
   return m.Next.IsZero() && m.Name == "" && m.Streak == 0 && len(m.Params) == 0
}

func (m *MetaAlg) ParamsStr() string {
   return strings.Join(m.Params, " ")
}

func (m *MetaAlg) String() (s string) {
   if !m.IsZero() {
      s = fmt.Sprintf("%s %d", m.NextStr(), m.Streak)

      if m.Name != "" {
         s += " " + m.Name
         if ps := m.ParamsStr(); ps != "" {
            s += " " + ps
         }
      }
   }

   return
}
