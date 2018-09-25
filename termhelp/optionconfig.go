package termhelp

import "errors"
import "fmt"
import "github.com/alanxoc3/concards/constring"
import "strconv"

type UsageMode int

const (
	VIEWMODE  UsageMode = iota
	EDITMODE            = iota
	PRINTMODE           = iota
)

type Config struct {
	// The various true or false options
	Review        bool
	Memorize      bool
	Done          bool
	NumberEnabled bool
	GroupsEnabled bool

	Help    bool
	Version bool
	Color   bool

	Usage UsageMode

	UpdateMode bool // Just writes files.
	Editor     string

	// The variable options passed in.
	Number      int
	Groups      map[string]bool
	GroupsSlice []string
	Files       []string

	Opts []*Option
}

// Parses through the given arguments and returns a generated config.
func ParseConfig(args []string) (*Config, error) {
	cfg := configInit()
	cfg.Opts = genOptions()

	argsNoProg := args[1:]
	curLen := 0
	var tmpArg string // Used in the for loop, for the options passed.

	pc := parseConfig{}

	for _, arg := range argsNoProg {
		curLen = len(arg)

		if curLen == 0 {
			// ERROR empty parameter
			return nil, errors.New("Error: You entered an empty parameter.")
		}

		if pc.waitForGroup { // PARSE GROUP STRINGS
			pc.waitForGroup = false
			lst := constring.StringToList(arg)
			if len(*lst) == 0 {
				return nil, fmt.Errorf("Error: You tried to pass an empty group list.")
			}

			for _, x := range *lst {
				if cfg.Groups[x] == false {
					cfg.Groups[x] = true
				} else { // ERROR same group
					return nil, errors.New("Error: You tried to pass the same group multiple times.")
				}
			}

		} else if pc.waitForNum { // PARSE NUMBER
			pc.waitForNum = false
			num, err := strconv.Atoi(arg)
			if err != nil {
				return nil, errors.New("Error: You didn't pass a number to the number option.")
			}
			cfg.Number = num
		} else if pc.waitForEditor { // PARSE STRING
			pc.waitForEditor = false
			cfg.Editor = arg
		} else {
			if arg[0] == '-' {
				if curLen == 1 {
					// ERROR, there is an argument with just a dash!
					return nil, errors.New("Error: You entered a dash with no options.")
				}

				if arg[1] == '-' { // Double Dash (Mario Kart?)
					if curLen == 2 {
						// ERROR, there is an argument with just two dashes!
						return nil, errors.New("Error: You entered two dashes with no options.")
					}
					tmpArg = arg[2:]
					err := executeCommand(&tmpArg, &pc, cfg)
					if err != nil {
						// ERROR, the command was not found.
						return nil, err
					}
				} else { // The Single Dash
					tmpArg = arg[1:]
					for i := 0; i < len(tmpArg); i++ {
						if pc.check() {
							return nil, errors.New("Error: You didn't provide a parameter for one of the options.")
						}
						err := executeAlias(tmpArg[i], &pc, cfg)
						if err != nil {
							// ERROR, the command was not found.
							return nil, err
						}
					}
				}
			} else { // This is a file!
				cfg.Files = append(cfg.Files, arg)
			}
		}
	}

	if pc.check() {
		return nil, errors.New("Error: You didn't provide a parameter for one of the options.")
	}

	return cfg, nil
}

// For debugging purposes.
func (cfg *Config) Print() {
	fmt.Printf("REV - MEM - DON - num - grp - hlp - ver\n")
	fmt.Printf("%t %t %t %t %t %t %t\n", cfg.Review,
		cfg.Memorize, cfg.Done, cfg.NumberEnabled, cfg.GroupsEnabled, cfg.Help,
		cfg.Version)
	fmt.Printf("VIE - EDI - PRI - UPD\n")
	fmt.Printf("%t %d\n", cfg.UpdateMode, cfg.Usage)

	fmt.Printf("ED: %s | NUM: %d | GRP %v | FIL %v | GRPSLC %v\n\n", cfg.Editor,
		cfg.Number, cfg.Groups, cfg.Files, cfg.GroupsSlice)
}

// Helpers...
func executeCommandWithNumber(num int, pc *parseConfig, cfg *Config) error {
	switch num {
	case REVIEW:
		cfg.Review = true
	case MEMORIZE:
		cfg.Memorize = true
	case DONE:
		cfg.Done = true
	case GROUPS:
		pc.waitForGroup = true
		cfg.GroupsEnabled = true
	case NUMBER:
		pc.waitForNum = true
		cfg.NumberEnabled = true
	case ONE:
		cfg.Number = 1
		cfg.NumberEnabled = true
	case EDIT:
		cfg.Usage = EDITMODE
	case PRINT:
		cfg.Usage = PRINTMODE
	case UPDATE:
		cfg.UpdateMode = true
	case HELP:
		cfg.Help = true
	case VERSION:
		cfg.Version = true
	case EDITOR:
		pc.waitForEditor = true
	default:
		// It doesn't exist here
		return errors.New("Error: You have an invalid command-line option.")
	}

	return nil
}

func executeCommand(cmd *string, pc *parseConfig, cfg *Config) error {
	num := optsFindCommand(cfg.Opts, cmd)
	return executeCommandWithNumber(num, pc, cfg)
}

func executeAlias(cmd byte, pc *parseConfig, cfg *Config) error {
	num := optsFindAlias(cfg.Opts, cmd)
	return executeCommandWithNumber(num, pc, cfg)
}

type parseConfig struct {
	waitForGroup, waitForNum, waitForEditor bool
}

// Set the defaults for the config
func configInit() *Config {
	// Everything besides these are set to false or 0
	var cfg Config
	cfg.Editor = ""
	cfg.Groups = make(map[string]bool)
	return &cfg
}

func (pc *parseConfig) check() bool {
	return pc.waitForGroup || pc.waitForNum || pc.waitForEditor
}
