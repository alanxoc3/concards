package internal_test

import (
	"testing"

	"github.com/alanxoc3/concards/internal"
	"github.com/stretchr/testify/assert"
)

func TestNewHash(t *testing.T) {
   h := internal.NewHash("a")
   assert.Equal(t, "a0000000000000000000000000000000", h.String())
}
