package card_test

import (
	"fmt"
	"testing"

	"github.com/alanxoc3/concards/internal/card"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCardErrNoQuestion(t *testing.T) {
	_, err := card.NewCards(".", "")
	assert.NotNil(t, err)
}

func TestNewCardErrNoFile(t *testing.T) {
	_, err := card.NewCards("", ".")
	assert.NotNil(t, err)
}

func TestOneNewCard(t *testing.T) {
	expectedHash := [16]byte{0x78, 0x14, 0xe1, 0xc8, 0xae, 0xa8, 0xa0, 0xc4, 0xca, 0x97, 0x0e, 0xdd, 0x85, 0x0b, 0x12, 0xf0}
	hashStr := fmt.Sprintf("%x", expectedHash)

	c, err := card.NewCards("hi", " hello  there  | \\: \\| \\@> \\<@  i'm  a  beard")

	require.Nil(t, err)
	require.Len(t, c, 1)
	require.Equal(t, 2, c[0].Len())
	assert.Equal(t, expectedHash, [16]byte(c[0].Hash()))
	assert.Equal(t, hashStr, c[0].Hash().String())
	assert.Equal(t, "hello there | \\: \\| \\@> \\<@ i'm a beard", c[0].String())
	assert.Equal(t, "hi", c[0].File())
	assert.Equal(t, "hello there", c[0].GetFactEsc(0))
	assert.Equal(t, ": | @> <@ i'm a beard", c[0].GetFactEsc(1))
	assert.True(t, c[0].HasAnswer())
}

func TestTwoNewCards(t *testing.T) {
	c, _ := card.NewCards(".", "question : answer")
	h := []string{"2abf30e888b3db27732dff3777687b74", "2fe8516253866fc0768f9ae6683d4bb5"}
	require.Len(t, c, len(h))
	assert.Equal(t, h[0], c[0].Hash().String())
	assert.Equal(t, h[1], c[1].Hash().String())
}

func TestColonNoSpace(t *testing.T) {
	c, _ := card.NewCards(".", "question:answer")
	h := []string{"2abf30e888b3db27732dff3777687b74", "2fe8516253866fc0768f9ae6683d4bb5"}
	require.Len(t, c, len(h))
	assert.Equal(t, h[0], c[0].Hash().String())
	assert.Equal(t, h[1], c[1].Hash().String())
}

func TestPipeNoSpace(t *testing.T) {
	c, _ := card.NewCards(".", "question|answer")
	h := "2abf30e888b3db27732dff3777687b74"
	assert.Equal(t, h, c[0].Hash().String())
}

func TestPipeBackslashNoSpace(t *testing.T) {
	c, _ := card.NewCards(".", "question \\|answer")
	h := "1a7d0274be552d6af1f0a998b988823f"
	assert.Equal(t, h, c[0].Hash().String())
	assert.Equal(t, "question \\|answer", c[0].String())
}

func TestBackslashSpace(t *testing.T) {
	c, _ := card.NewCards(".", "question \\  answer")
	require.Len(t, c, 1)
	assert.Equal(t, "question   answer", c[0].GetFactEsc(0))
}

func TestPipeThenColon(t *testing.T) {
   c, _ := card.NewCards(".", "question|:answer")
	require.Len(t, c, 2)
	assert.Equal(t, "question | answer", c[0].String())
	assert.Equal(t, "answer | question", c[1].String())
}

func TestColonThenPipe(t *testing.T) {
   c, _ := card.NewCards(".", "question:|answer")
	require.Len(t, c, 1)
	assert.Equal(t, "question | answer", c[0].String())
}

func TestDoubleBackslash(t *testing.T) {
   c, _ := card.NewCards(".", "question\\\\answer")
	require.Len(t, c, 1)
	assert.Equal(t, "question\\answer", c[0].GetFactEsc(0))
}

func TestBackslashSpaceNearEnd(t *testing.T) {
   c, _ := card.NewCards(".", "question answer\\ ")
	require.Len(t, c, 1)
	assert.Equal(t, "question answer ", c[0].GetFactEsc(0))
}

func TestBackslashSpaceVeryEnd(t *testing.T) {
   c, _ := card.NewCards(".", "question answer\\")
	require.Len(t, c, 1)
	assert.Equal(t, "question answer ", c[0].GetFactEsc(0))
}

func TestBackslashRandomLetter(t *testing.T) {
   c, _ := card.NewCards(".", "que\\stion answer")
	require.Len(t, c, 1)
	assert.Equal(t, "question answer", c[0].GetFactEsc(0))
	assert.Equal(t, "question answer", c[0].String())
}

func TestSpaceBetween(t *testing.T) {
   c, _ := card.NewCards(".", "question\\ answer")
	require.Len(t, c, 1)
	assert.Equal(t, "question answer", c[0].GetFactEsc(0))
	assert.Equal(t, "question answer", c[0].String())
}
