package meta

import (
	"fmt"
	"time"
)

type Result struct {
	meta
	Target bool
}

func NewResultFromPrediction(p *Prediction, target bool) *Result {
	r := &Result{
		meta:   p.meta,
		Target: target,
	}

	r.next = time.Now()
	if r.curr.IsZero() {
		r.curr = r.next
	}

	return r
}

func NewResultFromStrings(strs ...string) *Result {
	return &Result{
		meta:   *newMetaFromStrings(strs...),
		Target: getParam(strs, 6) == "1",
	}
}

func (r *Result) AnswerClassification() AnswerClassification {
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

func (r *Result) PredStreak() int {
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

func (r *Result) newCount(expecting bool, count int) int {
	if expecting == r.Target {
		count++
	}
	return count
}

func (r *Result) PredYesCount() int { return r.newCount(true, r.yesCount) }
func (r *Result) PredNoCount() int  { return r.newCount(false, r.noCount) }
func (r *Result) targetStr() string {
	if r.Target {
		return "1"
	} else {
		return "0"
	}
}

func (r *Result) String() string {
	return fmt.Sprintf("%s %s", r.meta.String(), r.targetStr())
}
