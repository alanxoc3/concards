package con

import "time"
import "fmt"
import "strings"

type Know uint16

const (
   NO  Know = iota
   IDK      = iota
   YES      = iota
)

type Meta struct {
   Next time.Time
   Streak int
   Name string
   Params []string
}

func NewAlg(ts string, streak string, name string, params []string) (m Meta) {
   m = Meta{}
   m.Name = name
   m.Next = timeOrNow(ts)
   m.Streak = intOrDefault(streak, 0)
   m.Params = params
   return
}

func (m Meta) Exec(input Know) (Meta, error) {
   switch m.Name {
      case "sm2": return sm2Exec(m, input), nil
      default: return m, fmt.Errorf("Algorithm doesn't exist")
   }
}

func (m Meta) NextStr() string {
   return m.Next.Format(time.RFC3339)
}

func (m Meta) ParamsStr() string {
   return strings.Join(m.Params, " ")
}

func (m Meta) String() (s string) {
   if !m.Next.IsZero() {
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
