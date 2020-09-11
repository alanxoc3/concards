package meta

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Meta interface {
	Hash() [16]byte
	Next() time.Time
	Curr() time.Time
	YesCount() int
	NoCount() int
	Streak() int

	HashStr() string
	String() string

	IsNew() bool
	IsZero() bool
}

type meta struct {
	hash     [16]byte
	next     time.Time
	curr     time.Time
	yesCount int
	noCount  int
	streak   int
}

func boundInt(num int, min int, max int) int {
	if min > max {
		panic("Logic error. Please report!")
	}

	if num > max {
		return max
	}
	if num < min {
		return min
	}
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
	if len(str)%2 == 1 {
		str += "0"
	}

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

func newMetaFromStrings(strs ...string) *meta {
	return newMetaBase(
		hashOrZero(getParam(strs, 0)),
		timeOrZero(getParam(strs, 1)),
		timeOrZero(getParam(strs, 2)),
		intOrZero(getParam(strs, 3)),
		intOrZero(getParam(strs, 4)),
		intOrZero(getParam(strs, 5)))
}

func newMetaBase(hash [16]byte, next time.Time, curr time.Time, yesCount int, noCount int, streak int) *meta {
	curr = curr.UTC()
	next = next.UTC()

	yesCount = boundInt(yesCount, 0, MaxYesNoStreak)
	noCount = boundInt(noCount, 0, MaxYesNoStreak)
	streak = boundInt(streak, -MaxYesNoStreak, MaxYesNoStreak)

	// Streak can't be larger than yes or no count.
	if streak > yesCount {
		yesCount = streak
	} else if streak < -noCount {
		noCount = -streak
	}

	return &meta{hash, next, curr, yesCount, noCount, streak}
}

func (m *meta) nextStr() string { return m.next.Format(time.RFC3339) }
func (m *meta) currStr() string { return m.curr.Format(time.RFC3339) }
func (m *meta) HashStr() string { return fmt.Sprintf("%x", m.hash) }
func (m *meta) String() string {
	return fmt.Sprintf("%s %s %s %d %d %d", m.HashStr(), m.nextStr(), m.currStr(), m.yesCount, m.noCount, m.streak)
}

func (m *meta) Hash() [16]byte  { return m.hash }
func (m *meta) Next() time.Time { return m.next }
func (m *meta) Curr() time.Time { return m.curr }
func (m *meta) YesCount() int   { return m.yesCount }
func (m *meta) NoCount() int    { return m.noCount }
func (m *meta) Streak() int     { return m.streak }

func (m *meta) IsNew() bool { return m.yesCount == 0 && m.noCount == 0 && m.streak == 0 }
func (m *meta) IsZero() bool {
	return bytes.Equal(m.hash[:], make([]byte, 16)) && m.next.IsZero() && m.curr.IsZero() && m.yesCount == 0 && m.noCount == 0 && m.streak == 0
}
