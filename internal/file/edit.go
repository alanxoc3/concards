package file

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/alanxoc3/concards/internal"
	"github.com/alanxoc3/concards/internal/card"
)

func EditCards(filename string, cfg *Config) (card.CardMap, error) {
	internal.AssertLogic(cfg != nil, "config was nil when passed to edit function")

	// Load the file with your favorite editor.
	cmd := exec.Command(cfg.Editor, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("Error: The editor returned an error code.")
	}

	return ReadCards(filename)
}
