package meta_test

import (
	"testing"
	"time"

	"github.com/alanxoc3/concards/meta"
	"github.com/stretchr/testify/assert"
)

func TestPredictionBasics(t *testing.T) {
	testMetaFuncs(t, func(strs ...string) meta.Meta {
		return meta.NewPredictionFromStrings(strs...)
	})
}

func TestPredictionString(t *testing.T) {
	p := meta.NewPredictionFromStrings("ff", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "12", "12", "12", "alg")
	assert.Equal(t, "ff000000000000000000000000000000 2020-01-01T00:00:00Z 2020-01-01T00:00:00Z 12 12 12 alg", p.String())
}

func TestPredictionStringEmpty(t *testing.T) {
	p := meta.NewPredictionFromStrings("ff", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "12", "12", "12")
	assert.Equal(t, "ff000000000000000000000000000000 2020-01-01T00:00:00Z 2020-01-01T00:00:00Z 12 12 12 ", p.String())
}

func TestPredictionDefault(t *testing.T) {
	p := meta.NewDefaultPrediction("ff", "hi")
	assert.Equal(t, "ff000000000000000000000000000000", p.HashStr())
	assert.True(t, p.Next().IsZero())
	assert.True(t, p.Curr().IsZero())
	assert.True(t, p.IsNew())
	assert.Equal(t, "hi", p.Name())
}

func TestPredictionIsNew(t *testing.T) {
	p := meta.NewDefaultPrediction("", "")
	assert.True(t, p.IsNew())
	assert.Zero(t, p.YesCount())
	assert.Zero(t, p.NoCount())
	assert.Zero(t, p.Streak())
}

func TestPredictionExecErr(t *testing.T) {
	p := meta.NewDefaultPrediction("", "hi")
	pp, err := p.Exec(true)
	assert.Nil(t, pp)
	assert.NotNil(t, err)
}

func TestPredictionExecErrIsNil(t *testing.T) {
	p := meta.NewDefaultPrediction("", "sm2")
	pp, err := p.Exec(true)
	assert.NotNil(t, pp)
	assert.Nil(t, err)
}

func TestPredictionExecHash(t *testing.T) {
	p := meta.NewDefaultPrediction("ff", "sm2")
	pp, _ := p.Exec(true)
	assert.Equal(t, "ff000000000000000000000000000000", pp.HashStr())
}

func TestPredictionExecCurr(t *testing.T) {
	p := meta.NewDefaultPrediction("", "sm2")
	tsOne := time.Now()
	pp, _ := p.Exec(true)
	tsTwo := time.Now()

	assert.True(t, tsOne.Equal(pp.Next()) || tsOne.Before(pp.Curr()))
	assert.True(t, tsTwo.Equal(pp.Next()) || tsTwo.After(pp.Curr()))
}

func TestPredictionExecYesCount(t *testing.T) {
	p := meta.NewDefaultPrediction("", "sm2")
	pp, _ := p.Exec(true)
	assert.Equal(t, 1, pp.YesCount())
	assert.Equal(t, 0, pp.NoCount())
	assert.Equal(t, 1, pp.Streak())
}

func TestPredictionExecNoCount(t *testing.T) {
	p := meta.NewDefaultPrediction("", "sm2")
	pp, _ := p.Exec(false)
	assert.Equal(t, 1, pp.NoCount())
	assert.Equal(t, 0, pp.YesCount())
	assert.Equal(t, -1, pp.Streak())
}

func TestPredictionExecName(t *testing.T) {
	p := meta.NewDefaultPrediction("", "sm2")
	pp, _ := p.Exec(false)
	assert.Equal(t, "sm2", pp.Name())
}
