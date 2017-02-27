package deck

import (
	"bufio"
	"os"
	"testing"

	"github.com/alanxoc3/concards-go/card"
	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	d, err := Open("sample_deck.txt")
	assert.NotNil(t, d, "returned nil instead of deck")
	assert.NoError(t, err, "returned error")
	assert.Equal(t, 3, len(d.Cards), "empty deck")
}

func TestSize(t *testing.T) {
	d := &Deck{}
	d.Cards = append(d.Cards, &card.Card{})
	assert.Equal(t, 1, d.Size(), "wrong size")
}

func TestNextBlock(t *testing.T) {
	f, _ := os.Open("sample_deck.txt")
	scanner := bufio.NewScanner(f)
	scanner.Scan()

	_, b := nextBlock(scanner, 0)
	assert.Equal(t, 1, b.end, "bad end line")
	assert.Equal(t, 1, len(b.lines), "empty block")

	blockCount := 1
	for {
		eof, _ := nextBlock(scanner, 0)
		blockCount++
		if eof {
			break
		}
	}

	assert.Equal(t, 5, blockCount, "wrong number of blocks")
}

func TestPullGroup(t *testing.T) {
	var b block
	b.lines = []string{"## Group1"}

	g := pullGroup(b)
	assert.NotEqual(t, "", g, "failed to pull group")
}
