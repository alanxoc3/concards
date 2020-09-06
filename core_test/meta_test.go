package card_test

import "github.com/stretchr/testify/assert"
import "github.com/alanxoc3/concards/core"
import "testing"

func TestNewMetaBaseIsZeroNoParams(t *testing.T) {
   m := core.NewMetaAlgFromStrings([]string{})
   assert.True(t, m.IsZero())

   assert.Empty(t, m.Hash)
   assert.True(t, m.Next.IsZero())
   assert.True(t, m.Curr.IsZero())
   assert.Zero(t, m.YesCount)
   assert.Zero(t, m.NoCount)
   assert.Zero(t, m.Streak)
}

func TestNewMetaBaseNotZero(t *testing.T) {
   m := core.NewMetaAlgFromStrings([]string{"hi"})
   assert.False(t, m.IsZero())
}

func TestNewMetaBaseIsZeroBadParams(t *testing.T) {
   m := core.NewMetaAlgFromStrings([]string{"", "date1", "date2", "bad", "boo", "beep"})
   assert.True(t, m.IsZero())
}

func TestNewMetaBase(t *testing.T) {
   m := core.NewMetaAlgFromStrings([]string{"", "date1", "date2", "bad", "boo", "beep"})
   assert.True(t, m.IsZero())
}
