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

type MetaBase {
	Next   time.Time
	Streak int
	Curr   time.Time
}

type MetaAlg struct {
   *MetaBase
	Params []string
	Name   string
}

type MetaHist struct {
   *MetaBase
   Target bool
}

func NewMetaBase(next string, streak string) *Meta {
	return &MetaBase{
      Next: timeOrNow(next),
		Streak: intOrDefault(streak, 0),
      Curr: time.Now(),
	}
}

func NewMetaAlg(next string, streak string, name string, params []string) *Meta {
	return &Meta{
      MetaBase: NewMetaBase(next, streak)
      Params: params,
		Name: name,
	}
}

func NewMetaHist(next string, streak string, bool target) *Meta {
	return &Meta{
      MetaBase: NewMetaBase(next, streak)
      Target: target,
	}
}

func NewDefaultMetaAlg(name string) *Meta {
	return &Meta{
      MetaBase: NewMetaBase(time.Now, 0),
      Params: []string{},
		Name: name,
	}
}

func (m *Meta) Exec(input bool) *Meta, error {
   nm := &Meta{}
   *nm = *m

   // Save the current time for logging & not saving the current time multiple times.
   nm.Next = time.Now()
   switch nm.Name {
      case "sm2": nm = sm2Exec(nm, input)
      default: return m, fmt.Errorf("Algorithm doesn't exist.")
   }

   // Streak Logic
   switch nm.GetAnswerCategory() {
      case YesWasYes: nm.Streak++
      case NoWasNo:   nm.Streak--
      default: nm.Streak=0
   }

   return nm, nil
}

func (m *Meta) GetAnswerCategory(input bool) AnswerCategory {
   if input {
      if m.Streak < 0 { return YesWasNo }
      else { return YesWasYes }
   } else {
      if m.Streak > 0 { return NoWasYes }
      else { return NoWasNo }
   }
}

func (m *Meta) IsZero() bool {
	return m.Next.IsZero() && m.Name == "" && m.Streak == 0 && len(m.Params) == 0
}

func (m *Meta) NextStr() string {
	return m.Next.Format(time.RFC3339)
}

func (m *Meta) ParamsStr() string {
	return strings.Join(m.Params, " ")
}

func (m *Meta) String() (s string) {
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
