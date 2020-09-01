package card_test

import "github.com/stretchr/testify/assert"
import "github.com/alanxoc3/concards/core"
import "testing"
import "time"

func TestNewMetaBase(t *testing.T) {
  assert := assert.New(t)
  ts := time.Now()
  m := core.NewMetaBase([]string{})

  assert.Empty(m.Hash)
  assert.False(ts.After(m.Next))
  assert.False(ts.After(m.Curr))
  assert.Zero(m.YesCount)
  assert.Zero(m.NoCount)
  assert.Zero(m.Streak)
}
