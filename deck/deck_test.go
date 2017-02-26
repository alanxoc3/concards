package deck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	d, err := Open("sample_deck.txt")
	assert.NotNil(t, d, "returned nil instead of deck")
	assert.NoError(t, err, "returned error")
	assert.True(t, len(d.Cards) > 0, "empty deck")
}
