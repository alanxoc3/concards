package deck

import (
	"time"

	"github.com/alanxoc3/concards/internal"
)

func pop(s []internal.Hash) ([]internal.Hash, internal.Hash) {
	internal.AssertLogic(len(s) > 0, "pop was called without checking state")
	return s[1:], s[0]
}

func (d *Deck) popReview() internal.Hash {
   stack, h := pop(d.reviewStack)
   d.reviewStack = stack
	return h
}

func (d *Deck) popFuture() internal.Hash {
   stack, h := pop(d.futureStack)
   d.futureStack = stack
	return h
}

func (d *Deck) canPopFutureStack(t time.Time) bool {
	if len(d.futureStack) > 0 {
		if val, ok := d.predictMap[d.futureStack[0]]; ok {
			return val.Next().Before(t)
		}
	}
	return false
}
