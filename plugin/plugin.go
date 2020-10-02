package plugin

import "github.com/alanxoc3/concards/internal/meta"

type Outcome = meta.Outcome

func MockOutcome(strs ...string) *Outcome {
   return meta.NewOutcomeFromStrings(strs...)
}
