package core

import "time"
import "strconv"

type MetaBase struct {
   Next     time.Time
   Curr     time.Time
   YesCount int
   NoCount  int
   Streak   int
}

func intOrDefault(str string, def int) int {
	if x, err := strconv.Atoi(str); err != nil {
		return def
	} else {
		return x
	}
}

func timeOrNow(str string) time.Time {
	if x, err := time.Parse(time.RFC3339, str); err != nil {
		return time.Now()
	} else {
		return x
	}
}

func NewMetaBaseFromStrings(next string, curr string, streak string) *MetaBase {
   return &MetaBase{
      Next: timeOrNow(next),
      Streak: intOrDefault(streak, 0),
      YesCount: 0,
      NoCount: 0,
      Curr: time.Now(),
   }
}

func (m *MetaBase) NextStr() string { return m.Next.Format(time.RFC3339) }
func (m *MetaBase) CurrStr() string { return m.Curr.Format(time.RFC3339) }
