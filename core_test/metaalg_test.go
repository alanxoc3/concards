package core_test

import "github.com/stretchr/testify/assert"
import "github.com/alanxoc3/concards/core"
import "testing"

func TestMetaAlgBasics(t *testing.T) {
   testMetaFuncs(t, func(strs ...string) core.Meta {
      return core.NewMetaAlgFromStrings(strs...)
   })
}

func TestMetaAlgString(t *testing.T) {
   m := core.NewMetaAlgFromStrings("ff000000000000000000000000000000", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "12", "12", "12", "alg")
   assert.Equal(t, "ff000000000000000000000000000000 2020-01-01T00:00:00Z 2020-01-01T00:00:00Z 12 12 12 alg", m.String())
}

func TestDefaultMetaAlg(t *testing.T) {
   m := core.NewMetaAlgFromStrings("ff000000000000000000000000000000", "2020-01-01T00:00:00Z", "2020-01-01T00:00:00Z", "12", "12", "12", "alg")
   assert.Equal(t, "ff000000000000000000000000000000 2020-01-01T00:00:00Z 2020-01-01T00:00:00Z 12 12 12 alg", m.String())
}
