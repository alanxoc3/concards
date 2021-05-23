package meta

import (
	"fmt"
	"math"
	"time"

	"github.com/alanxoc3/concards/internal"
)

type PredictMap map[internal.Hash]*Predict

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

func (p *Predict) Clone(h internal.Hash) *Predict {
	np := *p
	np.hash = h
	return &np
}

func (p *Predict) Exec(input bool, now time.Time) Predict {
	name := p.name

	algFunc, exists := algs[p.name]
	if !exists {
		algFunc = sm2Exec
		name = "sm2"
	}

	// Note that r.Next() has the current time.
	r := NewOutcomeFromPredict(p, now, input)
	next := r.Next().Add(time.Duration(math.Min(algFunc(*r), internal.MaxNextDuration)))

	return Predict{
		*newBase(
			r.Hash(),
			next,
			r.Next(),
			r.PredYesCount(),
			r.PredNoCount(),
			r.PredStreak(),
		), name,
	}
}

func (b *Predict) Name() string {
	return b.name
}

func (b *Predict) String() string {
	return fmt.Sprintf("%s %s", b.base.String(), b.name)
}
