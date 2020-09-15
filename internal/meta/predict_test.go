package meta_test

import (
	"testing"
	"time"

	"github.com/alanxoc3/concards/internal/meta"
	"github.com/stretchr/testify/assert"
)

func TestPredictBasics(t *testing.T) {
	testMetaFuncs(t, func(strs ...string) meta.Meta {
		return meta.NewPredictFromStrings(strs...)
	})
}

func TestPredictString(t *testing.T) {
	p := meta.NewPredictFromStrings("ff", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "12", "12", "12", "alg")
	assert.Equal(t, "ff000000000000000000000000000000 2020-01-01T00:00:00Z 2020-01-01T00:00:00Z 12 12 12 alg", p.String())
}

func TestPredictStringEmpty(t *testing.T) {
	p := meta.NewPredictFromStrings("ff", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "12", "12", "12")
	assert.Equal(t, "ff000000000000000000000000000000 2020-01-01T00:00:00Z 2020-01-01T00:00:00Z 12 12 12 ", p.String())
}

func TestPredictDefault(t *testing.T) {
	p := meta.NewDefaultPredict("ff", "hi")
	assert.Equal(t, "ff000000000000000000000000000000", p.Hash().String())
	assert.True(t, p.Next().IsZero())
	assert.True(t, p.Curr().IsZero())
	assert.Zero(t, p.Total())
	assert.Equal(t, "hi", p.Name())
}

func TestPredictIsNew(t *testing.T) {
	p := meta.NewDefaultPredict("", "")
	assert.Zero(t, p.Total())
	assert.Zero(t, p.YesCount())
	assert.Zero(t, p.NoCount())
	assert.Zero(t, p.Streak())
}

func TestPredictExecErr(t *testing.T) {
	p := meta.NewDefaultPredict("", "hi")
	pp, err := p.Exec(true)
	assert.Nil(t, pp)
	assert.NotNil(t, err)
}

func TestPredictExecErrIsNil(t *testing.T) {
	p := meta.NewDefaultPredict("", "sm2")
	pp, err := p.Exec(true)
	assert.NotNil(t, pp)
	assert.Nil(t, err)
}

func TestPredictExecHash(t *testing.T) {
	p := meta.NewDefaultPredict("ff", "sm2")
	pp, _ := p.Exec(true)
	assert.Equal(t, "ff000000000000000000000000000000", pp.Hash().String())
}

func TestPredictExecCurr(t *testing.T) {
	p := meta.NewDefaultPredict("", "sm2")
	tsOne := time.Now()
	pp, _ := p.Exec(true)
	tsTwo := time.Now()

	assert.True(t, tsOne.Equal(pp.Next()) || tsOne.Before(pp.Curr()))
	assert.True(t, tsTwo.Equal(pp.Next()) || tsTwo.After(pp.Curr()))
}

func TestPredictExecYesCount(t *testing.T) {
	p := meta.NewDefaultPredict("", "sm2")
	pp, _ := p.Exec(true)
	assert.Equal(t, 1, pp.YesCount())
	assert.Equal(t, 0, pp.NoCount())
	assert.Equal(t, 1, pp.Streak())
}

func TestPredictExecNoCount(t *testing.T) {
	p := meta.NewDefaultPredict("", "sm2")
	pp, _ := p.Exec(false)
	assert.Equal(t, 1, pp.NoCount())
	assert.Equal(t, 0, pp.YesCount())
	assert.Equal(t, -1, pp.Streak())
}

func TestPredictExecName(t *testing.T) {
	p := meta.NewDefaultPredict("", "sm2")
	pp, _ := p.Exec(false)
	assert.Equal(t, "sm2", pp.Name())
}