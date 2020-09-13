package plugin_test

import (
	"testing"

	"github.com/alanxoc3/concards/plugin"
	"github.com/stretchr/testify/assert"
)

func TestApi(t *testing.T) {
   assert.Equal(t, plugin.YesWasYes, plugin.MockOutcome("", "", "", "", "", "1", "1").AnswerClassification())
   assert.Equal(t, plugin.YesWasNo, plugin.MockOutcome("", "", "", "", "", "-1", "1").AnswerClassification())
   assert.Equal(t, plugin.NoWasYes, plugin.MockOutcome("", "", "", "", "", "1", "0").AnswerClassification())
   assert.Equal(t, plugin.NoWasNo, plugin.MockOutcome("", "", "", "", "", "-1", "0").AnswerClassification())
}
