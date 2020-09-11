package meta_test

import (
	"testing"
	"time"

	"github.com/alanxoc3/concards/meta"
	"github.com/stretchr/testify/assert"
)

func TestMetaAlgBasics(t *testing.T) {
	testMetaFuncs(t, func(strs ...string) meta.Meta {
		return meta.NewPredictionFromStrings(strs...)
	})
}

func TestMetaAlgString(t *testing.T) {
	ma := meta.NewPredictionFromStrings("ff", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "12", "12", "12", "alg")
	assert.Equal(t, "ff000000000000000000000000000000 2020-01-01T00:00:00Z 2020-01-01T00:00:00Z 12 12 12 alg", ma.String())
}

func TestMetaAlgStringEmpty(t *testing.T) {
	ma := meta.NewPredictionFromStrings("ff", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "12", "12", "12")
	assert.Equal(t, "ff000000000000000000000000000000 2020-01-01T00:00:00Z 2020-01-01T00:00:00Z 12 12 12 ", ma.String())
}

func TestMetaAlgDefault(t *testing.T) {
	ma := meta.NewDefaultPrediction("ff", "hi")
	assert.Equal(t, "ff000000000000000000000000000000", ma.HashStr())
	assert.True(t, ma.Next().IsZero())
	assert.True(t, ma.Curr().IsZero())
	assert.True(t, ma.IsNew())
	assert.Equal(t, "hi", ma.Name())
}

func TestMetaAlgIsNew(t *testing.T) {
	ma := meta.NewDefaultPrediction("", "")
	assert.True(t, ma.IsNew())
	assert.Zero(t, ma.YesCount())
	assert.Zero(t, ma.NoCount())
	assert.Zero(t, ma.Streak())
}

func TestMetaAlgExecErr(t *testing.T) {
	ma := meta.NewDefaultPrediction("", "hi")
	mma, err := ma.Exec(true)
	assert.Nil(t, mma)
	assert.NotNil(t, err)
}

func TestMetaAlgExecErrIsNil(t *testing.T) {
	ma := meta.NewDefaultPrediction("", "sm2")
	mma, err := ma.Exec(true)
	assert.NotNil(t, mma)
	assert.Nil(t, err)
}

func TestMetaAlgExecHash(t *testing.T) {
	ma := meta.NewDefaultPrediction("ff", "sm2")
	mma, _ := ma.Exec(true)
	assert.Equal(t, "ff000000000000000000000000000000", mma.HashStr())
}

func TestMetaAlgExecCurr(t *testing.T) {
	ma := meta.NewDefaultPrediction("", "sm2")
	tsOne := time.Now()
	mma, _ := ma.Exec(true)
	tsTwo := time.Now()

	assert.True(t, tsOne.Equal(mma.Next()) || tsOne.Before(mma.Curr()))
	assert.True(t, tsTwo.Equal(mma.Next()) || tsTwo.After(mma.Curr()))
}

func TestMetaAlgExecYesCount(t *testing.T) {
	ma := meta.NewDefaultPrediction("", "sm2")
	mma, _ := ma.Exec(true)
	assert.Equal(t, 1, mma.YesCount())
	assert.Equal(t, 0, mma.NoCount())
	assert.Equal(t, 1, mma.Streak())
}

func TestMetaAlgExecNoCount(t *testing.T) {
	ma := meta.NewDefaultPrediction("", "sm2")
	mma, _ := ma.Exec(false)
	assert.Equal(t, 1, mma.NoCount())
	assert.Equal(t, 0, mma.YesCount())
	assert.Equal(t, -1, mma.Streak())
}

func TestMetaAlgExecName(t *testing.T) {
	ma := meta.NewDefaultPrediction("", "sm2")
	mma, _ := ma.Exec(false)
	assert.Equal(t, "sm2", mma.Name())
}
