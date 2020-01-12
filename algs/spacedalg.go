package algs

import (
	"strconv"
	"time"
)

type Know uint16

const (
	NO  Know = iota
	IDK      = iota
	YES      = iota
)

type Exec func(s SpaceAlg, input Know) SpaceAlg

type SpaceAlg struct {
   Next time.Time
   Name string
   Streak int
   Params []string
   Exec Exec
}

func intOrDefault(str string, def int) int {
   if x, err := strconv.Atoi(str); err != nil {
      return def
   } else {
      return x
   }
}

func floatOrDefault(str string, def float32) float32 {
   if x, err := strconv.ParseFloat(str, 32); err != nil {
      return def
   } else {
      return float32(x)
   }
}

func timeOrToday(str string) time.Time {
   if x, err := time.Parse(time.RFC3339, str); err != nil {
      return time.Now()
   } else {
      return x
   }
}

func New(words []string) (space SpaceAlg) {
	space = SpaceAlg{}
   space.Streak = 0
   space.Next = time.Now()
   space.Params = []string{};

   for i, v := range words {
      switch i {
         case 0: space.Name = v
         case 1: space.Next = timeOrToday(v)
         case 2: space.Streak = intOrDefault(v, 0)
         case 3: space.Params = words[2:]
      }
   }

   switch space.Name {
      case "sm2": space.Exec = sm2Exec()
      default:
         space.Name = "sm2"
         space.Exec = sm2Exec()
   }

	return
}
