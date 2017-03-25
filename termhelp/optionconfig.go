package gui

import "errors"
import "fmt"
import "github.com/alanxoc3/concards-go/constring"
import "strconv"

type Config struct {
	// The various true or false options
	review        bool
	memorize      bool
	done          bool
	numberEnabled bool
	groupsEnabled bool

	help    bool
	version bool
	color   bool

	viewMode   bool
	editMode   bool
	printMode  bool
	updateMode bool

	mainScreen bool
	write      bool
	editor     string
	defalg     string

	// The variable options passed in.
	number int
	groups map[string]bool
	files  []string

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
			return nil, errors.New("You entered an empty parameter.")
		}

		if pc.waitForGroup { // PARSE GROUP STRINGS
			pc.waitForGroup = false
			lst := constring.StringToList(arg)
			for _, x := range *lst {
				if cfg.groups[x] == false {
					cfg.groups[x] = true
				} else { // ERROR same group
					return nil, errors.New("You tried to pass the same group multiple times.")
				}
			}

		} else if pc.waitForNum { // PARSE NUMBER
			pc.waitForNum = false
			num, err := strconv.Atoi(arg)
			if err != nil {
				return nil, errors.New("You didn't pass a number to the number option.")
			}
			cfg.number = num
		} else if pc.waitForDefAlg { // PARSE STRING
			pc.waitForDefAlg = false
			cfg.defalg = arg
		} else if pc.waitForEditor { // PARSE STRING
			pc.waitForEditor = false
			cfg.editor = arg
		} else {
			if arg[0] == '-' {
				if curLen == 1 {
					// ERROR, there is an argument with just a dash!
					return nil, errors.New("You entered a dash with no options.")
				}

				if arg[1] == '-' { // Double Dash (Mario Kart?)
					if curLen == 2 {
						// ERROR, there is an argument with just two dashes!
						return nil, errors.New("You entered two dashes with no options.")
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
							return nil, errors.New("You didn't provide a parameter for one of the options.")
						}
						err := executeAlias(tmpArg[i], &pc, cfg)
						if err != nil {
							// ERROR, the command was not found.
							return nil, err
						}
					}
				}
			} else { // This is a file!
				cfg.files = append(cfg.files, arg)
			}
		}
	}

	if pc.check() {
		return nil, errors.New("You didn't provide a parameter for one of the options.")
	}

	return cfg, nil
}

// For debugging purposes.
func (cfg *Config) Print() {
	fmt.Printf("REV - MEM - DON - num - grp - hlp - ver - col - scr - wri\n")
	fmt.Printf("%t %t %t %t %t %t %t %t %t %t\n", cfg.review,
		cfg.memorize, cfg.done, cfg.numberEnabled, cfg.groupsEnabled, cfg.help,
		cfg.version, cfg.color, cfg.mainScreen, cfg.write)

	fmt.Printf("ED: %s | DEF: %s | NUM: %d | GRP %v | FIL %v\n\n", cfg.editor,
		cfg.defalg, cfg.number, cfg.groups, cfg.files)
}

// Helpers...
func executeCommandWithNumber(num int, pc *parseConfig, cfg *Config) error {
	switch num {
	case REVIEW:
		cfg.review = true
	case MEMORIZE:
		cfg.memorize = true
	case DONE:
		cfg.done = true
	case GROUPS:
		pc.waitForGroup = true
		cfg.groupsEnabled = true
	case NUMBER:
		pc.waitForNum = true
	case ONE:
		cfg.number = 1
		cfg.numberEnabled = true
	case EDIT:
		cfg.editMode = true
	case PRINT:
		cfg.printMode = true
	case UPDATE:
		cfg.updateMode = true
	case HELP:
		cfg.help = true
	case VERSION:
		cfg.version = true
	case COLOR:
		cfg.color = true
	case NOMAIN:
		cfg.mainScreen = false
	case NOWRITE:
		cfg.write = false
	case EDITOR:
		pc.waitForEditor = true
	case DEFALG:
		pc.waitForDefAlg = true
	default:
		// It doesn't exist here
		return errors.New("You have an invalid command-line option.")
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
	waitForGroup, waitForNum, waitForDefAlg, waitForEditor bool
}

// Set the defaults for the config
func configInit() *Config {
	// Everything besides these are set to false or 0
	var cfg Config
	cfg.mainScreen = true
	cfg.viewMode = true
	cfg.editor = "$EDITOR"
	cfg.defalg = "SM2"
	cfg.groups = make(map[string]bool)
	return &cfg
}

func (pc *parseConfig) check() bool {
	return pc.waitForGroup || pc.waitForNum || pc.waitForDefAlg || pc.waitForEditor
}
