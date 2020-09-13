package meta

import (
	"fmt"
	"time"
)

type Outcome struct {
	base
	Target bool
}

func NewOutcomeFromPredict(p *Predict, target bool) *Outcome {
	r := &Outcome{
		base:   p.base,
		Target: target,
	}

	r.next = time.Now()
	if r.curr.IsZero() {
		r.curr = r.next
	}

	return r
}

func NewOutcomeFromStrings(strs ...string) *Outcome {
	return &Outcome{
		base:   *newMetaFromStrings(strs...),
		Target: getParam(strs, 6) == "1",
	}
}

func (r *Outcome) AnswerClassification() AnswerClassification {
	if r.Target {
		if r.streak < 0 {
			return YesWasNo
		} else {
			return YesWasYes
		}
	} else {
		if r.streak > 0 {
			return NoWasYes
		} else {
			return NoWasNo
		}
	}
}

func (r *Outcome) PredStreak() int {
	// Streak Logic
	streak := r.streak
	switch r.AnswerClassification() {
	case YesWasYes:
		streak++
	case NoWasNo:
		streak--
	default:
		streak = 0
	}
	return streak
}

func (r *Outcome) newCount(expecting bool, count int) int {
	if expecting == r.Target {
		count++
	}
	return count
}

func (r *Outcome) PredYesCount() int { return r.newCount(true, r.yesCount) }
func (r *Outcome) PredNoCount() int  { return r.newCount(false, r.noCount) }
func (r *Outcome) targetStr() string {
	if r.Target {
		return "1"
	} else {
		return "0"
	}
}

func (r *Outcome) RKey() RKey {
   return RKey{ r.Hash(), r.Total() }
}

func (r *Outcome) String() string {
	return fmt.Sprintf("%s %s", r.base.String(), r.targetStr())
}
