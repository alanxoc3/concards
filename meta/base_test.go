package meta_test

import (
	"testing"
	"time"

	"github.com/alanxoc3/concards/meta"
	"github.com/stretchr/testify/assert"
)

type metaCreate func(...string) meta.Meta
type metaTestFunc func(*testing.T, metaCreate)

func assertZero(t *testing.T, m meta.Meta) {
   assert.Zero(t, m.Hash())
   assert.True(t, m.Next().IsZero())
   assert.True(t, m.Curr().IsZero())
   assert.Zero(t, m.YesCount())
   assert.Zero(t, m.NoCount())
   assert.Zero(t, m.Streak())
}

func testMetaFuncs(t *testing.T, createFunc metaCreate) {
	for k, v := range metaTestFuncs {
		t.Run(k, func(t *testing.T) { v(t, createFunc) })
	}
}

var metaTestFuncs = map[string]metaTestFunc{
	"Hash": func(t *testing.T, cf metaCreate) {
		m := cf("ff")
		assert.Equal(t, "ff000000000000000000000000000000", m.Hash().String())
		assert.NotZero(t, m.Hash())
	},

   "HashEqual": func(t *testing.T, cf metaCreate) {
      p1 := cf("ff")
      p2 := cf("ff00")
      assert.True(t, p1.Hash() == p2.Hash())
   },

	"HashTooLong": func(t *testing.T, cf metaCreate) {
		m := cf("ff0000000000000000000000000000ff11")
		assert.Equal(t, "ff0000000000000000000000000000ff", m.Hash().String())
		assert.NotZero(t, m.Hash())
	},

	"HashOdd": func(t *testing.T, cf metaCreate) {
		m := cf("fab1e")
		assert.Equal(t, "fab1e000000000000000000000000000", m.Hash().String())
		assert.NotZero(t, m.Hash())
	},

	"Next": func(t *testing.T, cf metaCreate) {
		m := cf("", "2020-01-01T00:00:00Z")
		assert.NotZero(t, m)
		assert.Equal(t, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), m.Next())
	},

	"NoCount": func(t *testing.T, cf metaCreate) {
		m := cf("", "", "", "", "33")
		assert.NotZero(t, m)
		assert.Equal(t, 33, m.NoCount())
	},

	"StreakYes": func(t *testing.T, cf metaCreate) {
		m := cf("", "", "", "", "", "33")
		assert.NotZero(t, m)
		assert.Equal(t, 33, m.Streak())
		assert.Equal(t, 33, m.YesCount())
		assert.Equal(t, 0, m.NoCount())
	},

	"StreakNo": func(t *testing.T, cf metaCreate) {
		m := cf("", "", "", "", "", "-33")
		assert.NotZero(t, m)
		assert.Equal(t, -33, m.Streak())
		assert.Equal(t, 0, m.YesCount())
		assert.Equal(t, 33, m.NoCount())
	},

	"NoCountMax": func(t *testing.T, cf metaCreate) {
		m := cf("", "", "", "", "2000000001")
		assert.NotZero(t, m)
		assert.Equal(t, 536870912, m.NoCount())
	},

	"YesCountMax": func(t *testing.T, cf metaCreate) {
		m := cf("", "", "", "2000000001")
		assert.NotZero(t, m)
		assert.Equal(t, 536870912, m.YesCount())
	},

	"StreakMin": func(t *testing.T, cf metaCreate) {
		m := cf("", "", "", "", "", "-2000000001")
		assert.NotZero(t, m)
		assert.Equal(t, -536870912, m.Streak())
	},

	"StreakMax": func(t *testing.T, cf metaCreate) {
		m := cf("", "", "", "", "", "2000000001")
		assert.NotZero(t, m)
		assert.Equal(t, 536870912, m.Streak())
		assert.Equal(t, 536870912, m.Total())
	},

	"TotalMax": func(t *testing.T, cf metaCreate) {
		m := cf("", "", "", "2000000031", "2002000001", "2000040001")
		assert.Equal(t, 1073741824, m.Total())
	},

	"NextUTC": func(t *testing.T, cf metaCreate) {
		m := cf("", "2020-01-01T00:00:00-06:00")
		assert.NotZero(t, m)
		assert.Equal(t, time.Date(2020, 1, 1, 6, 0, 0, 0, time.UTC), m.Next())
	},

	"CurrUTC": func(t *testing.T, cf metaCreate) {
		m := cf("", "", "2020-01-01T00:00:00-06:00")
		assert.NotZero(t, m)
		assert.Equal(t, time.Date(2020, 1, 1, 6, 0, 0, 0, time.UTC), m.Curr())
	},

	"IsZeroNoParams": func(t *testing.T, cf metaCreate) {
		m := cf()
      assertZero(t, m)
	},

	"IsZeroBadParams": func(t *testing.T, cf metaCreate) {
		m := cf("", "date1", "date2", "bad", "boo", "beep")
		assertZero(t, m)
	},
	"Curr": func(t *testing.T, cf metaCreate) {
		m := cf("", "", "2020-01-01T00:00:00Z")
		assert.Equal(t, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), m.Curr())
	},

	"YesCount": func(t *testing.T, cf metaCreate) {
		m := cf("", "", "", "33")
		assert.Equal(t, 33, m.YesCount())
	},

	"NoCountMin": func(t *testing.T, cf metaCreate) {
		m := cf("", "", "", "", "-2000000001")
		assertZero(t, m)
	},

	"YesCountMin": func(t *testing.T, cf metaCreate) {
		m := cf("", "", "", "-2000000001")
		assertZero(t, m)
		assert.Equal(t, 0, m.YesCount())
	},
}
