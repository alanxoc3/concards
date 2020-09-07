package core

import "fmt"
import "time"
import "encoding/hex"
import "bytes"
import "strconv"

type metaBase struct {
   Hash     [16]byte
   Next     time.Time
   Curr     time.Time
   YesCount int
   NoCount  int
   Streak   int
}

func getParam(arr []string, i int) string {
   if i < len(arr) {
      return arr[i]
   } else {
      return ""
   }
}

func hashOrZero(str string) (hash [16]byte) {
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

func newMetaBase(strs []string) *metaBase {
   mb := &metaBase{}

   mb.Hash     = hashOrZero(getParam(strs, 0))
   mb.Next     = timeOrZero(getParam(strs, 1))
   mb.Curr     = timeOrZero(getParam(strs, 2))
   mb.YesCount = intOrZero(getParam(strs, 3))
   mb.NoCount  = intOrZero(getParam(strs, 4))
   mb.Streak   = intOrZero(getParam(strs, 5))

   return mb
}

func (m *metaBase) NextStr() string { return m.Next.Format(time.RFC3339) }
func (m *metaBase) CurrStr() string { return m.Curr.Format(time.RFC3339) }
func (m *metaBase) HashStr() string { return fmt.Sprintf("%x", m.Hash) }
func (m *metaBase) String()  string { return fmt.Sprintf("%s %s %s %d %d %d", m.HashStr(), m.NextStr(), m.CurrStr(), m.YesCount, m.NoCount, m.Streak) }

func (m *metaBase) isZero() bool { return bytes.Equal(m.Hash[:], make([]byte, 16)) && m.Next.IsZero() && m.Curr.IsZero() && m.YesCount == 0 && m.NoCount == 0 && m.Streak == 0 }
