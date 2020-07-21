package core

import "time"
import "fmt"
import "strings"

type Know uint16

var YearsToAddForKnown = 1000

const (
	NO  Know = iota
	IDK      = iota
	YES      = iota
	KNOW     = iota
)

type Meta struct {
	Next   time.Time
	Streak int
	Name   string
	Params []string
}

func NewMeta(ts string, streak string, name string, params []string) *Meta {
	return &Meta{
		Next:   timeOrNow(ts),
		Streak: intOrDefault(streak, 0),
		Name:   name,
		Params: params,
	}
}

func NewDefaultMeta(name string) *Meta {
	return &Meta{
		Next:   time.Now(),
		Streak: 0,
		Name:   name,
		Params: []string{},
	}
}

func (m *Meta) IsZero() bool {
	return m.Next.IsZero() && m.Name == "" && m.Streak == 0 && len(m.Params) == 0
}

func (m *Meta) getKnowIt() *Meta {
   newMeta := *m
   newMeta.Next = m.Next.AddDate(YearsToAddForKnown, 0, 0)
   return &newMeta
}

func (m *Meta) Exec(input Know) (*Meta, error) {
   if input == KNOW {
      return m.getKnowIt(), nil
   } else {
      switch m.Name {
      case "sm2":
         return sm2Exec(*m, input), nil
      default:
         return m, fmt.Errorf("Algorithm doesn't exist")
      }
   }
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
