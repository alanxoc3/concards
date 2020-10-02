package meta

import (
	"fmt"
	"time"
)

type OutcomeMap map[Key]*Outcome

type Outcome struct {
	base
	target bool
}

func NewOutcomeFromPredict(p *Predict, now time.Time, target bool) *Outcome {
	r := &Outcome{
		base:   p.base,
		target: target,
	}

	r.next = now.UTC()

	// For new cards, the outcome makes more sense to be instant.
	if r.curr.IsZero() {
		r.curr = now.UTC()
	}

	return r
}

func NewOutcomeFromStrings(strs ...string) *Outcome {
	return &Outcome{
		base:   *newMetaFromStrings(strs...),
		target: getParam(strs, 6) == "1",
	}
}

func (r *Outcome) PredYesCount() int { return r.newCount(true, r.yesCount) }
func (r *Outcome) PredNoCount() int  { return r.newCount(false, r.noCount) }
func (r *Outcome) PredStreak() int {
   if r.target && r.streak >= 0 {
      return r.streak + 1
   } else if !r.target && r.streak <= 0 {
      return r.streak - 1
   } else {
      return 0
   }
}

func (r *Outcome) Target() bool { return r.target }

func (r *Outcome) String() string {
	return fmt.Sprintf("%s %s", r.base.String(), r.targetStr())
}

func (r *Outcome) targetStr() string {
	if r.target {
		return "1"
	} else {
		return "0"
	}
}

func (r *Outcome) newCount(expecting bool, count int) int {
	if expecting == r.target {
		count++
	}
	return count
}
