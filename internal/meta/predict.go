package meta

import (
	"fmt"
	"math"
	"time"

	"github.com/alanxoc3/concards/internal"
)

type Predict struct {
	base
	name string
}

func NewPredictFromStrings(strs ...string) *Predict {
	return &Predict{
		base: *newMetaFromStrings(strs...),
		name: getParam(strs, 6),
	}
}

func NewDefaultPredict(hash internal.Hash, name string) *Predict {
	return &Predict{
		base: *newMetaFromStrings([]string{hash.String()}...),
		name: name,
	}
}

func (p *Predict) Exec(input bool) (*Predict, error) {
	// Note that r.Next() has the current time.
	r := NewOutcomeFromPredict(p, input)

	var next time.Time
	if algFunc, exists := algs[p.name]; exists {
		next = r.Next().Add(time.Duration(math.Min(algFunc(*r), internal.MaxNextDuration)))
	} else {
		return nil, fmt.Errorf("Algorithm doesn't exist.")
	}

	return &Predict{
		*newBase(
			r.Hash(),
			next,
			r.Next(),
			r.PredYesCount(),
			r.PredNoCount(),
			r.PredStreak(),
		), p.Name(),
	}, nil
}

func (b *Predict) Name() string {
	return b.name
}

func (b *Predict) String() string {
	return fmt.Sprintf("%s %s", b.base.String(), b.name)
}
