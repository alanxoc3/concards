package core_test

import "github.com/stretchr/testify/assert"
import "github.com/alanxoc3/concards/core"
import "testing"

func TestMetaHistBasics(t *testing.T) {
   testMetaFuncs(t, func(strs ...string) core.Meta {
      return core.NewMetaHistFromStrings(strs...)
   })
}

func TestMetaHistStringBadInput(t *testing.T) {
   m := core.NewMetaHistFromStrings("ff000000000000000000000000000000", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "12", "12", "12", "alg")
   assert.Equal(t, "ff000000000000000000000000000000 2020-01-01T00:00:00Z 2020-01-01T00:00:00Z 12 12 12 0", m.String())
}

func TestMetaHistStringTrue(t *testing.T) {
   m := core.NewMetaHistFromStrings("ff000000000000000000000000000000", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "12", "12", "12", "1")
   assert.Equal(t, "ff000000000000000000000000000000 2020-01-01T00:00:00Z 2020-01-01T00:00:00Z 12 12 12 1", m.String())
}

func TestMetaHistStringFalse(t *testing.T) {
   m := core.NewMetaHistFromStrings("ff000000000000000000000000000000", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "12", "12", "12")
   assert.Equal(t, "ff000000000000000000000000000000 2020-01-01T00:00:00Z 2020-01-01T00:00:00Z 12 12 12 0", m.String())
}
