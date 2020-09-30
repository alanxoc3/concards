package plugin_test

import (
	"testing"

	"github.com/alanxoc3/concards/plugin"
	"github.com/stretchr/testify/assert"
)

func TestApi(t *testing.T) {
   assert.True(t, plugin.MockOutcome("", "", "", "", "", "", "1").Target())
}
