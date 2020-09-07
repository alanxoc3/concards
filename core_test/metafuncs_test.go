package core_test

import "github.com/stretchr/testify/assert"
import "github.com/alanxoc3/concards/core"
import "testing"
import "time"

type metaCreate func(...string) core.Meta;
type metaTestFunc func(*testing.T, metaCreate);

func testMetaFuncs(t *testing.T, createFunc metaCreate) {
   for k, v := range metaTestFuncs {
      t.Run(k, func(t *testing.T) { v(t, createFunc) })
   }
}

var metaTestFuncs = map[string]metaTestFunc{
   "Hash": func(t *testing.T, cf metaCreate) {
      m := cf("ff")
      assert.Equal(t, "ff000000000000000000000000000000", m.HashStr())
      assert.NotZero(t, m.Hash())
   },

   "Next": func(t *testing.T, cf metaCreate) {
      m := cf("", "2020-01-01T00:00:00Z")
      assert.False(t, m.IsZero())
      assert.Equal(t, time.Date(2020,1,1,0,0,0,0,time.UTC), m.Next())
   },

   "NoCount": func(t *testing.T, cf metaCreate) {
      m := cf("", "", "", "", "33")
      assert.False(t, m.IsZero())
      assert.Equal(t, 33, m.NoCount())
   },

   "StreakYes": func(t *testing.T, cf metaCreate) {
      m := cf("", "", "", "", "", "33")
      assert.False(t, m.IsZero())
      assert.Equal(t, 33, m.Streak())
      assert.Equal(t, 33, m.YesCount())
      assert.Equal(t, 0, m.NoCount())
   },

   "StreakNo": func(t *testing.T, cf metaCreate) {
      m := cf("", "", "", "", "", "-33")
      assert.False(t, m.IsZero())
      assert.Equal(t, -33, m.Streak())
      assert.Equal(t, 0, m.YesCount())
      assert.Equal(t, 33, m.NoCount())
   },

   "NoCountMax": func(t *testing.T, cf metaCreate) {
      m := cf("", "", "", "", "2000000001")
      assert.False(t, m.IsZero())
      assert.Equal(t, 1073741824, m.NoCount())
   },

   "YesCountMax": func(t *testing.T, cf metaCreate) {
      m := cf("", "", "", "2000000001")
      assert.False(t, m.IsZero())
      assert.Equal(t, 1073741824, m.YesCount())
   },

   "StreakMin": func(t *testing.T, cf metaCreate) {
      m := cf("", "", "", "", "", "-2000000001")
      assert.False(t, m.IsZero())
      assert.Equal(t, -1073741824, m.Streak())
   },

   "StreakMax": func(t *testing.T, cf metaCreate) {
      m := cf("", "", "", "", "", "2000000001")
      assert.False(t, m.IsZero())
      assert.Equal(t, 1073741824, m.Streak())
   },

   "NextUTC": func(t *testing.T, cf metaCreate) {
      m := cf("", "2020-01-01T00:00:00-06:00")
      assert.False(t, m.IsZero())
      assert.Equal(t, time.Date(2020,1,1,6,0,0,0,time.UTC), m.Next())
   },

   "CurrUTC": func(t *testing.T, cf metaCreate) {
      m := cf("", "", "2020-01-01T00:00:00-06:00")
      assert.False(t, m.IsZero())
      assert.Equal(t, time.Date(2020,1,1,6,0,0,0,time.UTC), m.Curr())
   },

   "IsZeroNoParams": func(t *testing.T, cf metaCreate) {
      m := cf()
      assert.True(t, m.IsZero())

      assert.Zero(t, m.Hash())
      assert.True(t, m.Next().IsZero())
      assert.True(t, m.Curr().IsZero())
      assert.Zero(t, m.YesCount())
      assert.Zero(t, m.NoCount())
      assert.Zero(t, m.Streak())
   },

   "IsZeroBadParams": func(t *testing.T, cf metaCreate) {
      m := cf("", "date1", "date2", "bad", "boo", "beep")
      assert.True(t, m.IsZero())
   },
   "Curr": func(t *testing.T, cf metaCreate) {
      m := cf("", "", "2020-01-01T00:00:00Z")
      assert.Equal(t, time.Date(2020,1,1,0,0,0,0,time.UTC), m.Curr())
   },

   "YesCount": func(t *testing.T, cf metaCreate) {
      m := cf("", "", "", "33")
      assert.Equal(t, 33, m.YesCount())
   },

   "NoCountMin": func(t *testing.T, cf metaCreate) {
      m := cf("", "", "", "", "-2000000001")
      assert.Zero(t, 0, m)
      assert.Equal(t, 0, m.NoCount())
   },

   "YesCountMin": func(t *testing.T, cf metaCreate) {
      m := cf("", "", "", "-2000000001")
      assert.Zero(t, 0, m)
      assert.Equal(t, 0, m.YesCount())
   },
}
