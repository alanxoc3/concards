package termhelp

import "fmt"

// Returns nil as Config if the program finished from this method. That would
// happend with an error, help, or version.
func ValidateAndParseConfig(args []string) (*Config, error) {
	cfg, err := ParseConfig(args)

	if err != nil {
		return nil, err
	}

	if !cfg.Help && !cfg.Version {
		if len(cfg.Files) == 0 {
			return nil, fmt.Errorf("Error: You didn't provide any files to review from.")
		}

		// None for all!
		if !cfg.Review && !cfg.Memorize && !cfg.Done {
			cfg.Review = true
			cfg.Memorize = true
			cfg.Done = false
		}

		// Make the group slice :).
		for k, v := range cfg.Groups {
			if v == true {
				cfg.GroupsSlice = append(cfg.GroupsSlice, k)
			}
		}

		if cfg.NumberEnabled && cfg.Number <= 0 {
			return nil, fmt.Errorf("Error: The number you passed was too small.")
		}

	} else if cfg.Help {
		fmt.Println(Help())
		return nil, nil
	} else if cfg.Version {
		fmt.Println(Version())
		return nil, nil
	}

	return cfg, err
}
