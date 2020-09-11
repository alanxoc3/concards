package _test

import (
	"testing"

	"github.com/alanxoc3/concards/meta"
	"github.com/stretchr/testify/assert"
)

func TestMetaHistBasics(t *testing.T) {
	testMetaFuncs(t, func(strs ...string) meta.Meta {
		return meta.NewResultFromStrings(strs...)
	})
}

func TestMetaHistStringBadInput(t *testing.T) {
	mh := meta.NewResultFromStrings("ff000000000000000000000000000000", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "12", "12", "12", "alg")
	assert.Equal(t, "ff000000000000000000000000000000 2020-01-01T00:00:00Z 2020-01-01T00:00:00Z 12 12 12 0", mh.String())
}

func TestMetaHistStringTrue(t *testing.T) {
	mh := meta.NewResultFromStrings("ff000000000000000000000000000000", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "12", "12", "12", "1")
	assert.Equal(t, "ff000000000000000000000000000000 2020-01-01T00:00:00Z 2020-01-01T00:00:00Z 12 12 12 1", mh.String())
}

func TestMetaHistStringFalse(t *testing.T) {
	mh := meta.NewResultFromStrings("ff000000000000000000000000000000", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "12", "12", "12")
	assert.Equal(t, "ff000000000000000000000000000000 2020-01-01T00:00:00Z 2020-01-01T00:00:00Z 12 12 12 0", mh.String())
}
