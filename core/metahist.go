package core

import "time"
import "fmt"
import "strings"

type AnswerCategory uint8

const (
   YesWasYes AnswerCategory = 1 << iota
   YesWasNo
   NoWasYes
   NoWasNo
)

type MetaHist struct {
   Hash string
   *MetaBase
   Target bool
}

func NewMetaHistFromMetaAlg(hash string, ma *MetaAlg, target bool) *MetaHist {
   return &MetaHist{
      Hash: hash,
      MetaBase: NewMetaBase(time.Now(), ma.Curr, ma.Streak)
      Target: target,
   }
}

func (m *MetaHist) GetAnswerCategory() AnswerCategory {
   if m.Target {
      if m.Streak < 0 { return YesWasNo }
      else { return YesWasYes }
   } else {
      if m.Streak > 0 { return NoWasYes }
      else { return NoWasNo }
   }
}
