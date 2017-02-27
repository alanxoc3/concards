package card

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCard(t *testing.T) {
	c := Card{}
	assert.NotNil(t, c)
}

func TestNew(t *testing.T) {
	blocks := [][]string{
		[]string{
			"Question String",
			"\tAnswerString1",
			"\tAnswerString2",
		},
		[]string{
			"Question String",
			"Question String2",
			"\tAnswerString1",
		},
		[]string{
			"Question String",
			"Question String2",
			"\tAnswerString1",
			"Question String3",
		},
		[]string{
			"Question String",
			"\tAnswerString1",
			"~~ Metadata",
		},
		[]string{
			"Question String",
			"\tAnswerString1",
			"~~ Metadata",
			"Junk",
		},
	}

	successMap := []bool{true, true, false, true, false}

	for i, block := range blocks {
		_, err := New(block)
		if successMap[i] {
			assert.NoError(t, err, "failed parsing card")
		} else {
			assert.Error(t, err, "failed parsing card")
		}
	}
}
