package card

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCard(t *testing.T) {
	c := Card{}
	assert.NotNil(t, c)
}
