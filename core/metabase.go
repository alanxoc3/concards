package core

import "fmt"
import "time"
import "encoding/hex"
import "bytes"
import "strconv"

type Meta interface {
   NextStr() string
   CurrStr() string
   HashStr() string
   String()  string

   Hash() [16]byte
   Next() time.Time
   Curr() time.Time
   YesCount() int
   NoCount() int
   Streak() int

   IsNew() bool
   IsZero() bool
}

// About a billion.
const YesNoStreakMax = 1<<30

type metaBase struct {
   hash     [16]byte
   next     time.Time
   curr     time.Time
   yesCount int
   noCount  int
   streak   int
}

func boundInt(num int, min int, max int) int {
   if min > max { panic("Logic error. Please report!") }

   if num > max { return max }
   if num < min { return min }
   return num
}

func getParam(arr []string, i int) string {
   if i < len(arr) {
      return arr[i]
   } else {
      return ""
   }
}

func hashOrZero(str string) (hash [16]byte) {
   if len(str) % 2 == 1 { str += "0" }

	if x, err := hex.DecodeString(str); err == nil {
      copy(hash[:], x)
	}
   return hash
}

func intOrZero(str string) int {
	if x, err := strconv.Atoi(str); err != nil {
		return 0
	} else {
		return x
	}
}

func timeOrZero(str string) time.Time {
	if x, err := time.Parse(time.RFC3339, str); err != nil {
		return time.Time{}
	} else {
		return x
	}
}

func newMetaBaseFromStrings(strs ...string) *metaBase {
   return newMetaBase(
      hashOrZero(getParam(strs, 0)),
      timeOrZero(getParam(strs, 1)),
      timeOrZero(getParam(strs, 2)),
      intOrZero(getParam(strs, 3)),
      intOrZero(getParam(strs, 4)),
      intOrZero(getParam(strs, 5)))
}

func newMetaBase(hash [16]byte, next time.Time, curr time.Time, yesCount int, noCount int, streak int) *metaBase {
   curr = curr.UTC()
   next = next.UTC()

   yesCount = boundInt(yesCount, 0, YesNoStreakMax)
   noCount = boundInt(noCount,  0, YesNoStreakMax)
   streak = boundInt(streak, -YesNoStreakMax, YesNoStreakMax)

   // Streak can't be larger than yes or no count.
   if streak > yesCount {
      yesCount = streak
   } else if streak < -noCount {
      noCount = -streak
   }

   return &metaBase{ hash, next, curr, yesCount, noCount, streak }
}

func (m *metaBase) NextStr() string { return m.next.Format(time.RFC3339) }
func (m *metaBase) CurrStr() string { return m.curr.Format(time.RFC3339) }
func (m *metaBase) HashStr() string { return fmt.Sprintf("%x", m.hash) }
func (m *metaBase) String()  string { return fmt.Sprintf("%s %s %s %d %d %d", m.HashStr(), m.NextStr(), m.CurrStr(), m.yesCount, m.noCount, m.streak) }

func (m *metaBase) Hash() [16]byte { return m.hash }
func (m *metaBase) Next() time.Time { return m.next }
func (m *metaBase) Curr() time.Time { return m.curr }
func (m *metaBase) YesCount() int { return m.yesCount }
func (m *metaBase) NoCount() int { return m.noCount }
func (m *metaBase) Streak() int { return m.streak }

func (m *metaBase) IsNew() bool { return m.yesCount == 0 && m.noCount == 0 && m.streak == 0 }
func (m *metaBase) IsZero() bool { return bytes.Equal(m.hash[:], make([]byte, 16)) && m.next.IsZero() && m.curr.IsZero() && m.yesCount == 0 && m.noCount == 0 && m.streak == 0 }
