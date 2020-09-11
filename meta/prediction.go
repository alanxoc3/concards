package meta

import (
	"fmt"
	"math"
	"time"
)

type Prediction struct {
	meta
	name string
}

func NewPredictionFromStrings(strs ...string) *Prediction {
	return &Prediction{
		meta: *newMetaFromStrings(strs...),
		name: getParam(strs, 6),
	}
}

func NewDefaultPrediction(hash string, name string) *Prediction {
	return &Prediction{
		meta: *newMetaFromStrings([]string{hash}...),
		name: name,
	}
}

func (p *Prediction) Exec(input bool) (*Prediction, error) {
	// Note that r.Next() has the current time.
	r := NewResultFromPrediction(p, input)

	var next time.Time
	if algFunc, exists := algs[p.name]; exists {
		next = r.Next().Add(time.Duration(math.Min(algFunc(*r), MaxNextDuration)))
	} else {
		return nil, fmt.Errorf("Algorithm doesn't exist.")
	}

	return &Prediction{
		*newMetaBase(
			r.Hash(),
			next,
			r.Next(),
			r.PredYesCount(),
			r.PredNoCount(),
			r.PredStreak(),
		), p.Name(),
	}, nil
}

func (m *Prediction) Name() string {
	return m.name
}

func (m *Prediction) String() string {
	return fmt.Sprintf("%s %s", m.meta.String(), m.name)
}
