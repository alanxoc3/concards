package file_test

import (
	"strings"
	"testing"

	"github.com/alanxoc3/concards/internal/card"
	"github.com/alanxoc3/concards/internal/file"
	"github.com/alanxoc3/concards/internal/meta"
	"github.com/stretchr/testify/assert"
)

func TestReadCards(t *testing.T) {
	fstr := " hello ye as \n @> hi | hello <@ asoe @> yoyo man go <@"
	cards := file.ReadCardsFromReader(strings.NewReader(fstr), "file")
	hi, _ := card.NewCards("file", "hi | hello")
	yo, _ := card.NewCards("file", "yoyo man go")

	assert.Equal(t, hi[0], cards[0])
	assert.Equal(t, yo[0], cards[1])
}

func TestReadPredicts(t *testing.T) {
	fstr := `a 2020-02-01T00:00:00Z 2020-01-01T00:00:00Z 1 2 0 sm2
            b 2020-02-01T00:00:00Z 2020-01-01T00:00:00Z 2 1 0 sm2
            c 2020-02-01T00:00:00Z 2020-01-01T00:00:00Z 3 3 3 sm2`

	predicts := file.ReadPredictsFromReader(strings.NewReader(fstr))
	assert.Len(t, predicts, 3)
	assert.Equal(t, predicts[0], meta.NewPredictFromStrings("a", "2020-02-01T00:00:00Z", "2020-01-01T00:00:00Z", "1", "2", "0", "sm2"))
	assert.Equal(t, predicts[1], meta.NewPredictFromStrings("b", "2020-02-01T00:00:00Z", "2020-01-01T00:00:00Z", "2", "1", "0", "sm2"))
	assert.Equal(t, predicts[2], meta.NewPredictFromStrings("c", "2020-02-01T00:00:00Z", "2020-01-01T00:00:00Z", "3", "3", "3", "sm2"))
}
