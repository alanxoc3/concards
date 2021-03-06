package meta_test

import (
	"testing"
	"time"

	"github.com/alanxoc3/concards/internal"
	"github.com/alanxoc3/concards/internal/meta"
	"github.com/stretchr/testify/assert"
)

func newDefPred(h string, a string) *meta.Predict {
	return meta.NewDefaultPredict(internal.NewHash(h), a)
}

func TestPredictBasics(t *testing.T) {
	testMetaFuncs(t, func(strs ...string) meta.Meta {
		return meta.NewPredictFromStrings(strs...)
	})
}

func TestPredictString(t *testing.T) {
	p := meta.NewPredictFromStrings("ff", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "1", "1", "0", "alg")
	assert.Equal(t, "ff000000000000000000000000000000 2020-01-01T00:00:00Z 2020-01-01T00:00:00Z 1 1 0 alg", p.String())
}

func TestPredictStringEmpty(t *testing.T) {
	p := meta.NewPredictFromStrings("ff", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "1", "1", "0")
	assert.Equal(t, "ff000000000000000000000000000000 2020-01-01T00:00:00Z 2020-01-01T00:00:00Z 1 1 0 ", p.String())
}

func TestPredictDefault(t *testing.T) {
	p := newDefPred("ff", "hi")
	assert.Equal(t, "ff000000000000000000000000000000", p.Hash().String())
	assert.True(t, p.Next().IsZero())
	assert.True(t, p.Curr().IsZero())
	assert.Zero(t, p.Total())
	assert.Equal(t, "hi", p.Name())
}

func TestPredictIsNew(t *testing.T) {
	p := newDefPred("", "")
	assert.Zero(t, p.Total())
	assert.Zero(t, p.YesCount())
	assert.Zero(t, p.NoCount())
	assert.Zero(t, p.Streak())
}

func TestPredictExecErr(t *testing.T) {
	p := newDefPred("", "hi")
	pp := p.Exec(true, time.Now())
	assert.Equal(t, "sm2", pp.Name())
}

func TestPredictExecErrIsNil(t *testing.T) {
	p := newDefPred("", "sm2")
	pp := p.Exec(true, time.Now())
	assert.NotNil(t, pp)
}

func TestPredictExecHash(t *testing.T) {
	p := newDefPred("ff", "sm2")
	pp := p.Exec(true, time.Now())
	assert.Equal(t, "ff000000000000000000000000000000", pp.Hash().String())
}

func TestPredictExecCurr(t *testing.T) {
	p := newDefPred("", "sm2")
	ts := time.Date(2021, 1, 1, 0, 0, 0, 2, time.UTC)
	pp := p.Exec(true, ts)

	assert.True(t, ts.Equal(pp.Curr()))
	assert.True(t, ts.Before(pp.Next()))
}

func TestPredictExecYesCount(t *testing.T) {
	p := newDefPred("", "sm2")
	pp := p.Exec(true, time.Now())
	assert.Equal(t, 1, pp.YesCount())
	assert.Equal(t, 0, pp.NoCount())
	assert.Equal(t, 1, pp.Streak())
}

func TestPredictExecNoCount(t *testing.T) {
	p := newDefPred("", "sm2")
	pp := p.Exec(false, time.Now())
	assert.Equal(t, 1, pp.NoCount())
	assert.Equal(t, 0, pp.YesCount())
	assert.Equal(t, -1, pp.Streak())
}

func TestPredictExecName(t *testing.T) {
	p := newDefPred("", "sm2")
	pp := p.Exec(false, time.Now())
	assert.Equal(t, "sm2", pp.Name())
}

func TestPredictClone(t *testing.T) {
	p := newDefPred("ff", "sm2")
	assert.Equal(t, p, p.Clone(internal.Hash{0xff}))
}
