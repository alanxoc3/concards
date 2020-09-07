package card_test

import "github.com/stretchr/testify/assert"
import "github.com/alanxoc3/concards/core"
import "testing"
import "time"

func TestNewMetaBaseIsZeroNoParams(t *testing.T) {
   m := core.NewMetaAlgFromStrings([]string{})
   assert.True(t, m.IsZero())

   assert.Equal(t, "00000000000000000000000000000000", m.HashStr())
   assert.True(t, m.Next.IsZero())
   assert.True(t, m.Curr.IsZero())
   assert.Zero(t, m.YesCount)
   assert.Zero(t, m.NoCount)
   assert.Zero(t, m.Streak)
}

func TestNewMetaBaseIsZeroBadParams(t *testing.T) {
   m := core.NewMetaAlgFromStrings([]string{"", "date1", "date2", "bad", "boo", "beep"})
   assert.True(t, m.IsZero())
}

func TestNewMetaBaseHash(t *testing.T) {
   m := core.NewMetaAlgFromStrings([]string{"ff"})
   assert.Equal(t, "ff000000000000000000000000000000", m.HashStr())
   assert.Equal(t, "ff000000000000000000000000000000 0001-01-01T00:00:00Z 0001-01-01T00:00:00Z 0 0 0 ", m.String())
   assert.False(t, m.IsZero())
}

func TestNewMetaBaseNext(t *testing.T) {
   m := core.NewMetaAlgFromStrings([]string{"", "2020-01-01T00:00:00Z"})
   assert.True(t, time.Date(2020,1,1,0,0,0,0,time.UTC).Equal(m.Next))
   assert.Equal(t, "00000000000000000000000000000000 2020-01-01T00:00:00Z 0001-01-01T00:00:00Z 0 0 0 ", m.String())
   assert.False(t, m.IsZero())
}

func TestNewMetaBaseCurr(t *testing.T) {
   m := core.NewMetaAlgFromStrings([]string{"", "", "2020-01-01T00:00:00Z"})
   assert.True(t, time.Date(2020,1,1,0,0,0,0,time.UTC).Equal(m.Curr))
   assert.Equal(t, "00000000000000000000000000000000 0001-01-01T00:00:00Z 2020-01-01T00:00:00Z 0 0 0 ", m.String())
   assert.False(t, m.IsZero())
}

func TestNewMetaBaseYesCount(t *testing.T) {
   m := core.NewMetaAlgFromStrings([]string{"", "", "", "33"})
   assert.Equal(t, 33, m.YesCount)
   assert.Equal(t, "00000000000000000000000000000000 0001-01-01T00:00:00Z 0001-01-01T00:00:00Z 33 0 0 ", m.String())
   assert.False(t, m.IsZero())
}

func TestNewMetaBaseNoCount(t *testing.T) {
   m := core.NewMetaAlgFromStrings([]string{"", "", "", "", "33"})
   assert.Equal(t, 33, m.NoCount)
   assert.Equal(t, "00000000000000000000000000000000 0001-01-01T00:00:00Z 0001-01-01T00:00:00Z 0 33 0 ", m.String())
   assert.False(t, m.IsZero())
}

func TestNewMetaBaseStreak(t *testing.T) {
   m := core.NewMetaAlgFromStrings([]string{"", "", "", "", "", "33"})
   assert.Equal(t, 33, m.Streak)
   assert.Equal(t, "00000000000000000000000000000000 0001-01-01T00:00:00Z 0001-01-01T00:00:00Z 0 0 33 ", m.String())
   assert.False(t, m.IsZero())
}

func TestNewMetaBaseName(t *testing.T) {
   m := core.NewMetaAlgFromStrings([]string{"", "", "", "", "", "", "alg"})
   assert.Equal(t, "alg", m.Name)
   assert.Equal(t, "00000000000000000000000000000000 0001-01-01T00:00:00Z 0001-01-01T00:00:00Z 0 0 0 alg", m.String())
   assert.False(t, m.IsZero())
}
