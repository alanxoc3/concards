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
	r := meta.NewOutcomeFromStrings("ff000000000000000000000000000000", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "2", "2", "1", "alg")
	assert.Equal(t, "ff000000000000000000000000000000 2020-01-01T00:00:00Z 2020-01-01T00:00:00Z 2 2 1 0", r.String())
}

func TestOutcomeStringTrue(t *testing.T) {
	r := meta.NewOutcomeFromStrings("ff000000000000000000000000000000", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "2", "2", "1", "1")
	assert.Equal(t, "ff000000000000000000000000000000 2020-01-01T00:00:00Z 2020-01-01T00:00:00Z 2 2 1 1", r.String())
}

func TestOutcomeStringFalse(t *testing.T) {
	r := meta.NewOutcomeFromStrings("ff000000000000000000000000000000", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "2", "2", "1")
	assert.Equal(t, "ff000000000000000000000000000000 2020-01-01T00:00:00Z 2020-01-01T00:00:00Z 2 2 1 0", r.String())
}

func TestOutcomeTargetTrue(t *testing.T) {
	r := meta.NewOutcomeFromStrings("ff000000000000000000000000000000", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "2", "2", "1", "1")
	assert.True(t, r.Target())
}

func TestOutcomeTargetFalse(t *testing.T) {
	r := meta.NewOutcomeFromStrings("ff000000000000000000000000000000", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "2", "2", "1", "0")
	assert.False(t, r.Target())
}
