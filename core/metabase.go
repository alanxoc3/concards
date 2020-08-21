package core

import "time"

type MetaBase struct {
   Next   time.Time
   Curr   time.Time
   Streak int
}

func NewMetaBaseFromStrings(next string, curr string, streak string) *MetaBase {
   return &MetaBase{
      Next: timeOrNow(next),
      Streak: intOrDefault(streak, 0),
      Curr: time.Now(),
   }
}

func (m *MetaBase) NextStr() string { return m.Next.Format(time.RFC3339) }
func (m *MetaBase) CurrStr() string { return m.Curr.Format(time.RFC3339) }
