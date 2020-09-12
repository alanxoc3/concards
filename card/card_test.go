package card_test

import (
	"encoding/hex"
	"testing"

	"github.com/alanxoc3/concards/card"
	"github.com/stretchr/testify/assert"
)

func TestNewCardErrNoQuestion(t *testing.T) {
	_, err := card.NewCards(".", "")
	assert.NotNil(t, err)
}

func TestNewCardErrNoFile(t *testing.T) {
	_, err := card.NewCards("", ".")
	assert.NotNil(t, err)
}

func TestNewCard(t *testing.T) {
	c, err := card.NewCards(".", " hello  there  |  i'm  a  beard")
   hashStr := "c435597dd9718c64b135087e944fd614"

   t.Run("NoErr", func(t *testing.T) {
      assert.Nil(t, err)
   })

   t.Run("Hash", func(t *testing.T) {
      expectedHash, _ := hex.DecodeString(hashStr)
      actualHash      := c[0].Hash()
      assert.Equal(t, expectedHash, actualHash[:])
   })

   t.Run("HashStr", func(t *testing.T) {
      assert.Equal(t, "c435597dd9718c64b135087e944fd614", c[0].HashStr())
   })

   t.Run("String", func(t *testing.T) {
      assert.Equal(t, "hello there | i'm a beard", c[0].String())
   })
}
