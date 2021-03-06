package meta

import (
	"fmt"
	"strconv"
	"time"

	"github.com/alanxoc3/concards/internal"
)

type Key struct {
	Hash  internal.Hash
	Total int
}

type Meta interface {
	Hash() internal.Hash
	Next() time.Time
	Curr() time.Time
	YesCount() int
	NoCount() int
	Streak() int

	String() string
	Total() int
	Key() Key
	IsZero() bool
}

type base struct {
	hash     internal.Hash
	next     time.Time
	curr     time.Time
	yesCount int
	noCount  int
	streak   int
}

func boundInt(num int, min int, max int) int {
	internal.AssertLogic(min <= max, "min was greater than max")

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

func newMetaFromStrings(strs ...string) *base {
	return newBase(
		internal.NewHash(getParam(strs, 0)),
		timeOrZero(getParam(strs, 1)),
		timeOrZero(getParam(strs, 2)),
		intOrZero(getParam(strs, 3)),
		intOrZero(getParam(strs, 4)),
		intOrZero(getParam(strs, 5)))
}

func newBase(hash internal.Hash, next time.Time, curr time.Time, yesCount int, noCount int, streak int) *base {
	curr = curr.UTC()
	next = next.UTC()

	yesCount = boundInt(yesCount, 0, internal.MaxYesNoStreak)
	noCount = boundInt(noCount, 0, internal.MaxYesNoStreak)

   if yesCount == 0 && noCount == 0 || yesCount == 1 && noCount == 1 || yesCount == 0 && streak > 0 || noCount == 0 && streak < 0 {
      streak = 0
   } else if noCount == 0 && yesCount > 0 {
      streak = yesCount
   } else if yesCount == 0 && noCount > 0 {
      streak = -noCount
   } else if streak >= yesCount {
      streak = yesCount - 1
   } else if -streak >= noCount {
      streak = -(noCount - 1)
   }

	// Streak can't be larger than yes or no count.
	if streak > yesCount {
		yesCount = streak
	} else if streak < -noCount {
		noCount = -streak
	}

	return &base{hash, next, curr, yesCount, noCount, streak}
}

func (b *base) nextStr() string { return b.next.Format(time.RFC3339) }
func (b *base) currStr() string { return b.curr.Format(time.RFC3339) }
func (b *base) String() string {
	return fmt.Sprintf("%s %s %s %d %d %d", b.Hash().String(), b.nextStr(), b.currStr(), b.yesCount, b.noCount, b.streak)
}

func (b *base) Hash() internal.Hash { return b.hash }
func (b *base) Next() time.Time     { return b.next }
func (b *base) Curr() time.Time     { return b.curr }
func (b *base) YesCount() int       { return b.yesCount }
func (b *base) NoCount() int        { return b.noCount }
func (b *base) Streak() int         { return b.streak }

func (b *base) Total() int {
	sum := b.yesCount + b.noCount
	internal.AssertLogic(sum >= 0, "total overflowed")
	return sum
}

func (b *base) Key() Key {
	return Key{b.Hash(), b.Total()}
}

func (b *base) IsZero() bool {
	return b.next.IsZero() && b.curr.IsZero() && b.Total() == 0
}
