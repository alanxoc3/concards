package card_test

import (
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
	c, err := card.NewCards("hi", " hello  there  | \\: \\| \\:> \\<:  i'm  a  beard")

	require.Nil(t, err)
	require.Len(t, c, 1)
	require.Equal(t, 2, c[0].Len())
	assert.Equal(t, "hello there | \\: \\| \\:> <\\: i'm a beard", c[0].String())
	assert.Equal(t, "hi", c[0].File())
	assert.Equal(t, "hello there", c[0].GetFactEsc(0))
	assert.Equal(t, ": | :> <: i'm a beard", c[0].GetFactEsc(1))
	assert.True(t, c[0].HasAnswer())
}

func TestOneSide(t *testing.T) {
	c, _ := card.NewCards(".", "a")
	require.Len(t, c, 1)
	assert.Equal(t, "a", c[0].GetFactEsc(0))
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
	require.Len(t, c, 1)
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

func TestBackslashVeryEnd(t *testing.T) {
	c, _ := card.NewCards(".", "question answer\\")
	require.Len(t, c, 1)
	assert.Equal(t, "question answer", c[0].GetFactEsc(0))
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

func TestSpaceInCurly(t *testing.T) {
   c, _ := card.NewCards(".", "a{ ha}b")
	require.Len(t, c, 1)
	assert.Equal(t, "a {}b | ha", c[0].String())
}

func TestExtraEndCurly(t *testing.T) {
   c, _ := card.NewCards(".", "a {}}")
	require.Len(t, c, 1)
	assert.Equal(t, "a {}\\}", c[0].String())
}

func TestBackslashHash(t *testing.T) {
   c, _ := card.NewCards(".", "heh #3")
	require.Len(t, c, 1)
	assert.Equal(t, "heh \\#3", c[0].String())
}

func TestClozeCrazySpaceExample(t *testing.T) {
   c, _ := card.NewCards(".", "hi{ hey{ you }{me{ inner }}}yo")
	require.Len(t, c, 4)
	assert.Equal(t, "hi {} yo | hey you me inner", c[0].String())
	assert.Equal(t, "hi hey {} me inner yo | you", c[1].String())
	assert.Equal(t, "hi hey you {} yo | me inner", c[2].String())
	assert.Equal(t, "hi hey you me {} yo | inner", c[3].String())
}

func TestClozeDoubleCurly(t *testing.T) {
   c, _ := card.NewCards(".", "a {{b}} c")
	require.Len(t, c, 2)
	assert.Equal(t, "a {} c | b", c[0].String())
	assert.Equal(t, "a {} c | b", c[1].String())
}

func TestClozeTwoOfThem(t *testing.T) {
   c, _ := card.NewCards(".", "{he}lp {him}")
	require.Len(t, c, 2)
	assert.Equal(t, "{}lp him | he", c[0].String())
	assert.Equal(t, "help {} | him", c[1].String())
}

func TestClozeDoubleCurlyMoreText(t *testing.T) {
   c, _ := card.NewCards(".", "a {hap{p}y} c")
	require.Len(t, c, 2)
	assert.Equal(t, "a {} c | happy", c[0].String())
	assert.Equal(t, "a hap{}y c | p", c[1].String())
}

func TestClozeGroupOne(t *testing.T) {
   c, _ := card.NewCards(".", "aaah #{he}lp #{me}")
	require.Len(t, c, 1)
	assert.Equal(t, "aaah {}lp {} | he | me", c[0].String())
}

func TestClozeOnlySpaces(t *testing.T) {
   c, _ := card.NewCards(".", "hi{   }yo")
	require.Len(t, c, 1)
	assert.Equal(t, "hi yo", c[0].String())
}

func TestClozeNestedOnlySpaces(t *testing.T) {
   c, _ := card.NewCards(".", "hi{ { } {   } }yo")
	require.Len(t, c, 3)
	assert.Equal(t, "hi yo", c[0].String())
	assert.Equal(t, "hi yo", c[1].String())
	assert.Equal(t, "hi yo", c[2].String())
}

func TestClozeGroupEmpty(t *testing.T) {
   c, _ := card.NewCards(".", "#{ }#{me}")
	require.Len(t, c, 1)
	assert.Equal(t, "{} | me", c[0].String())
}

func TestMultipleClozeGroups(t *testing.T) {
   c, _ := card.NewCards(".", "##{a} #{b} ##{c} #{d} {e}")
	require.Len(t, c, 3)
	assert.Equal(t, "a b c d {} | e", c[0].String())
	assert.Equal(t, "a {} c {} e | b | d", c[1].String())
	assert.Equal(t, "{} b {} d e | a | c", c[2].String())
}

func TestMultipleNestedClozeGroups(t *testing.T) {
   c, _ := card.NewCards(".", "##{a #{b}} ##{c #{d}} {e}")
	require.Len(t, c, 3)
	assert.Equal(t, "a b c d {} | e", c[0].String())
	assert.Equal(t, "a {} c {} e | b | d", c[1].String())
	assert.Equal(t, "{} {} e | a b | c d", c[2].String())
}

func TestNestedSameClozeGroup(t *testing.T) {
   c, _ := card.NewCards(".", "#{a #{b #{c}}} d")
	require.Len(t, c, 1)
	assert.Equal(t, "{} d | a b c", c[0].String())
}

func TestClozeGroup71(t *testing.T) {
   c, _ := card.NewCards(".", "#######################################################################{hell}o world")
	require.Len(t, c, 1)
	assert.Equal(t, "{}o world | hell", c[0].String())
}

func TestHashThenCloze(t *testing.T) {
   c, _ := card.NewCards(".", "# {hell}o")
   require.Len(t, c, 1)
   assert.Equal(t, "\\# {}o | hell", c[0].String())
}

func TestClozeColonGroup(t *testing.T) {
   c, _ := card.NewCards(".", "#{he:ll}o")
   require.Len(t, c, 1)
   assert.Equal(t, "{}o | hell", c[0].String())
}

func TestClozeColon(t *testing.T) {
   c, _ := card.NewCards(".", "{he:ll}o")
   require.Len(t, c, 2)
   assert.Equal(t, "{}llo | he", c[0].String())
   assert.Equal(t, "he{}o | ll", c[1].String())
}

func TestClozeMultipleColons(t *testing.T) {
   c, _ := card.NewCards(".", "{h:e:l:l}o")
   require.Len(t, c, 4)
   assert.Equal(t, "{}ello | h", c[0].String())
   assert.Equal(t, "h{}llo | e", c[1].String())
   assert.Equal(t, "he{}lo | l", c[2].String())
   assert.Equal(t, "hel{}o | l", c[3].String())
}

func TestClozeDoubleColon(t *testing.T) {
   c, _ := card.NewCards(".", "{h:e::l:l}o")
   require.Len(t, c, 2)
   assert.Equal(t, "\\{h\\:e | l\\:l\\}o", c[0].String())
   assert.Equal(t, "l\\:l\\}o | \\{h\\:e", c[1].String())
}

func TestClozeTripleColon(t *testing.T) {
   c, _ := card.NewCards(".", "{h:e:::l:l}o")
   require.Len(t, c, 2)
   assert.Equal(t, "\\{h\\:e | \\:l\\:l\\}o", c[0].String())
   assert.Equal(t, "\\:l\\:l\\}o | \\{h\\:e", c[1].String())
}

func TestClozeQuadrupleColon(t *testing.T) {
   c, _ := card.NewCards(".", "{h:e::::l:l}o")
   require.Len(t, c, 2)
   assert.Equal(t, "\\{h\\:e | l\\:l\\}o", c[0].String())
   assert.Equal(t, "l\\:l\\}o | \\{h\\:e", c[1].String())
}
