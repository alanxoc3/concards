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
	assert.Nil(t, err)
}

func TestNewCardErrNoFile(t *testing.T) {
	_, err := card.NewCards("", ".")
	assert.NotNil(t, err)
}

func TestOneNewCard(t *testing.T) {
	expectedHash := [16]byte{0xef, 0x88, 0xa5, 0x74, 0xb8, 0x35, 0xe4, 0xd8, 0x52, 0xd7, 0x36, 0x44, 0x52, 0x01, 0x38, 0x8f}
	hashStr := fmt.Sprintf("%x", expectedHash)

	c, err := card.NewCards("hi", " hello  there  | \\: \\| \\:> \\<:  i'm  a  beard")

	require.Nil(t, err)
	require.Len(t, c, 1)
	require.Equal(t, 2, c[0].Len())
	assert.Equal(t, expectedHash, [16]byte(c[0].Hash()))
	assert.Equal(t, hashStr, c[0].Hash().String())
	assert.Equal(t, "hello there | \\: \\| \\:\\> \\<\\: i'm a beard", c[0].String())
	assert.Equal(t, "hi", c[0].File())
	assert.Equal(t, "hello there", c[0].GetFactEsc(0))
	assert.Equal(t, ": | :> <: i'm a beard", c[0].GetFactEsc(1))
	assert.True(t, c[0].HasAnswer())
}

func TestTwoNewCards(t *testing.T) {
   c, _ := card.NewCards(".", "  |question :: answer")
	h := []string{"2abf30e888b3db27732dff3777687b74", "2fe8516253866fc0768f9ae6683d4bb5"}
	require.Len(t, c, len(h))
	assert.Equal(t, h[0], c[0].Hash().String())
	assert.Equal(t, h[1], c[1].Hash().String())
}

func TestColonNoSpace(t *testing.T) {
   c, _ := card.NewCards(".", "question::answer")
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
	c, _ := card.NewCards(".", "question \\ \nanswer")
	require.Len(t, c, 1)
	assert.Equal(t, "question   answer", c[0].GetFactEsc(0))
}

func TestPipeThenColon(t *testing.T) {
   c, _ := card.NewCards(".", "question|::answer")
	require.Len(t, c, 2)
	assert.Equal(t, "question | answer", c[0].String())
	assert.Equal(t, "answer | question", c[1].String())
}

func TestColonThenPipe(t *testing.T) {
   c, _ := card.NewCards(".", "question::|answer")
	require.Len(t, c, 2)
	assert.Equal(t, "question | answer", c[0].String())
	assert.Equal(t, "answer | question", c[1].String())
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
   c, _ := card.NewCards(".", "question\\  answer")
	require.Len(t, c, 1)
	assert.Equal(t, "question  answer", c[0].GetFactEsc(0))
	assert.Equal(t, "question\\  answer", c[0].String())
}

func TestColonEscaped(t *testing.T) {
   c, _ := card.NewCards(".", "q:a")
	require.Len(t, c, 1)
   assert.Equal(t, "q:a", c[0].GetFactEsc(0))
   assert.Equal(t, "q\\:a", c[0].String())
}

func TestThreeReversibleSides(t *testing.T) {
   c, _ := card.NewCards(".", "z | y :: a | b :: p")
	require.Len(t, c, 5)
	assert.Equal(t, "z | a | b | p", c[0].String())
	assert.Equal(t, "y | a | b | p", c[1].String())
	assert.Equal(t, "a | z | y | p", c[2].String())
	assert.Equal(t, "b | z | y | p", c[3].String())
	assert.Equal(t, "p | z | y | a | b", c[4].String())
}

func TestColonAtBeginning(t *testing.T) {
   c, _ := card.NewCards(".", ":: a | b")
	require.Len(t, c, 1)
	assert.Equal(t, "a | b", c[0].String())
}
