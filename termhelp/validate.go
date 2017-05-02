package termhelp

// After the config is parsed, this may check for possible errors before
// actually parsing the files.
func ValidateAndParseConfig(args []string) (*Config, error) {
	cfg, err := ParseConfig(args)

	if err != nil {
		return nil, err
	}

	return cfg, err
}
