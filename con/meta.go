package con

import "time"
import "fmt"

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
