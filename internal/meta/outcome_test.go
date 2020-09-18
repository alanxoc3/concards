package meta_test

import (
	"testing"

	"github.com/alanxoc3/concards/internal/meta"
	"github.com/stretchr/testify/assert"
)

func TestOutcomeBasics(t *testing.T) {
	testMetaFuncs(t, func(strs ...string) meta.Meta {
		return meta.NewOutcomeFromStrings(strs...)
	})
}

func TestOutcomeStringBadInput(t *testing.T) {
	r := meta.NewOutcomeFromStrings("ff000000000000000000000000000000", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "12", "12", "12", "alg")
	assert.Equal(t, "ff000000000000000000000000000000 2020-01-01T00:00:00Z 2020-01-01T00:00:00Z 12 12 12 0", r.String())
}

func TestOutcomeStringTrue(t *testing.T) {
	r := meta.NewOutcomeFromStrings("ff000000000000000000000000000000", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "12", "12", "12", "1")
	assert.Equal(t, "ff000000000000000000000000000000 2020-01-01T00:00:00Z 2020-01-01T00:00:00Z 12 12 12 1", r.String())
}

func TestOutcomeStringFalse(t *testing.T) {
	r := meta.NewOutcomeFromStrings("ff000000000000000000000000000000", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "12", "12", "12")
	assert.Equal(t, "ff000000000000000000000000000000 2020-01-01T00:00:00Z 2020-01-01T00:00:00Z 12 12 12 0", r.String())
}

func TestOutcomeTargetTrue(t *testing.T) {
	r := meta.NewOutcomeFromStrings("ff000000000000000000000000000000", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "12", "12", "12", "1")
	assert.True(t, r.Target())
}

func TestOutcomeTargetFalse(t *testing.T) {
	r := meta.NewOutcomeFromStrings("ff000000000000000000000000000000", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "12", "12", "12", "0")
	assert.False(t, r.Target())
}

func TestOutcomeAnswerClassification(t *testing.T) {
	tests := []struct {
		expect       meta.AnswerClassification
		expectStreak int
		streak       string
		answer       string
	}{
		{meta.YesWasYes, 2, "1", "1"},
		{meta.YesWasNo, 0, "-1", "1"},
		{meta.NoWasYes, 0, "1", "0"},
		{meta.NoWasNo, -2, "-1", "0"},
		{meta.YesWasYes, 1, "0", "1"},
		{meta.NoWasNo, -1, "0", "0"},
	}

	for _, v := range tests {
		r := meta.NewOutcomeFromStrings("", "", "", "", "", v.streak, v.answer)
		assert.Equal(t, v.expect, r.AnswerClassification())
		assert.Equal(t, v.expectStreak, r.PredStreak())
	}
}
