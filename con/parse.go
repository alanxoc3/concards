package con

import (
	"strconv"
	"time"
)

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

func timeOrNow(str string) time.Time {
   if x, err := time.Parse(time.RFC3339, str); err != nil {
      return time.Now()
   } else {
      return x
   }
}
