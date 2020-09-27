package stack_test

import (
	"testing"
	"time"

	"github.com/alanxoc3/concards/internal"
	"github.com/alanxoc3/concards/internal/stack"
	"github.com/stretchr/testify/assert"
)

var ONE_DATE time.Time = time.Date(1, 1, 1, 0, 0, 0, 1, time.UTC)
var TWO_DATE time.Time = time.Date(1, 1, 1, 0, 0, 0, 2, time.UTC)

func TestTime(t *testing.T) {
	s := stack.NewStack(ONE_DATE)
	assert.Equal(t, ONE_DATE, s.Time())
}

func TestSetTime(t *testing.T) {
	s := stack.NewStack(ONE_DATE)
	s.SetTime(TWO_DATE)
	assert.Equal(t, TWO_DATE, s.Time())
}

func TestEmpty(t *testing.T) {
	s := stack.NewStack(ONE_DATE)
	assert.Nil(t, s.Top())
	assert.Empty(t, s.List())
	assert.Zero(t, s.ReviewLen())
	assert.Zero(t, s.FutureLen())
	assert.False(t, s.Update(internal.NewHash(""), TWO_DATE))
}

func TestInsert(t *testing.T) {
	s := stack.NewStack(ONE_DATE)
   h := internal.NewHash("")
   s.Insert(h, ONE_DATE)
	assert.Len(t, s.List(), 1)
}

func TestInsertSameFuture(t *testing.T) {
	s := stack.NewStack(TWO_DATE)
   h := internal.NewHash("")
   s.Insert(h, TWO_DATE)
	assert.Equal(t, 0, s.ReviewLen())
	assert.Equal(t, 1, s.FutureLen())
}

func TestInsertDifferentFuture(t *testing.T) {
	s := stack.NewStack(ONE_DATE)
   h := internal.NewHash("")
   s.Insert(h, TWO_DATE)
	assert.Equal(t, 0, s.ReviewLen())
	assert.Equal(t, 1, s.FutureLen())
}

func TestInsertDifferentReview(t *testing.T) {
	s := stack.NewStack(TWO_DATE)
   h := internal.NewHash("")
   s.Insert(h, ONE_DATE)
	assert.Equal(t, h, *s.Top())
	assert.Equal(t, 1, s.ReviewLen())
	assert.Equal(t, 0, s.FutureLen())
}

func TestPop(t *testing.T) {
	s := stack.NewStack(TWO_DATE)
   s.Insert(internal.NewHash("ff"), ONE_DATE)
   s.Pop()
	assert.Len(t, s.List(), 0)
}

func TestClone(t *testing.T) {
	s1 := stack.NewStack(ONE_DATE)
	s2 := stack.NewStack(TWO_DATE)
   s1.Clone(s2)

	assert.Equal(t, s2, s1)
}

func TestListInsertionOrder(t *testing.T) {
	s := stack.NewStack(TWO_DATE)
   f := internal.NewHash("f")
   e := internal.NewHash("e")
   d := internal.NewHash("d")

   s.Insert(f, ONE_DATE)
   s.Insert(e, TWO_DATE)
   s.Insert(d, ONE_DATE)

   assert.Equal(t, []internal.Hash{f, e, d}, s.List())
}

func TestUpdate(t *testing.T) {
	s := stack.NewStack(TWO_DATE)
   f := internal.NewHash("f")
   e := internal.NewHash("e")
   d := internal.NewHash("d")

   s.Insert(f, TWO_DATE)
   s.Insert(e, TWO_DATE)
   s.Insert(d, TWO_DATE)

   s.Update(f, ONE_DATE)
   s.Update(d, ONE_DATE)

   assert.Equal(t, 2, s.ReviewLen())
   assert.Equal(t, 1, s.FutureLen())
   assert.Equal(t, f, *s.Top())
}
