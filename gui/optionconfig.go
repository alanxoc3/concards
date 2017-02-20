package gui

import "errors"
import "fmt"

type Config struct {
	// The various true or false options
	review   bool
	memorize bool
	done     bool
	numberEnabled bool
	groupsEnabled bool

	help         bool
	version      bool
	color        bool

	viewMode   bool
	editMode   bool
	printMode  bool
	updateMode bool

	mainScreen   bool
	write bool
	editor string
	defalg string


	// The variable options passed in.
	number int
	groups map[string]bool
	files []string

	Opts []*Option
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
	return &cfg
}

// Parses through the given arguments and returns a generated config.
func ParseConfig(args []string) (*Config, error) {
	cfg := configInit()
	cfg.Opts = GenOptions()

	argsNoProg := args[1:]
	curLen := 0
	var tmpArg string // Used in the for loop, for the options passed.

	pc := parseConfig{}

	for _,arg := range argsNoProg {
		curLen = len(arg)

		if curLen == 0 {
			// ERROR empty parameter
			return nil, errors.New("You entered an empty parameter.")
		}

		if pc.waitForGroup {
			pc.waitForGroup = false
			// PARSE GROUPS AND ADD THEM

		} else if pc.waitForNum {
			pc.waitForNum = false
			// PARSE NUMBER AND CHANGE THE NUMBER

		} else if pc.waitForDefAlg {
			pc.waitForDefAlg = false
			// PARSE THE DEFAULT_ALGORITHM

		} else if pc.waitForEditor {
			pc.waitForEditor = false
			// PARSE FOR THE EDITOR

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
				} else { // The Single Dash
					tmpArg = arg[1:]
				}
				err := executeCommand(&tmpArg,&pc,cfg)
				if err != nil {
					// ERROR, the command was not found.
					return nil, err
				}
			} else { // This is a file!
				cfg.files = append(cfg.files, arg)
			}
		}
	}

	if pc.waitForGroup || pc.waitForNum || pc.waitForDefAlg || pc.waitForEditor {
		return nil, errors.New("You didn't provide a parameter for one of the options.")
	}

	return cfg, nil
}

func executeCommandWithNumber(num int, pc *parseConfig, cfg *Config) error {
	switch num {
		case REVIEW:     cfg.review       = true
		case MEMORIZE:   cfg.memorize     = true
		case DONE:       cfg.done         = true
		case GROUPS:     pc.waitForGroup  = true
		                 cfg.groupsEnabled = true
		case NUMBER:     pc.waitForNum    = true
		case ONE:        cfg.number       = 1
		                 cfg.numberEnabled = true
		case EDIT:       cfg.editMode     = true
		case PRINT:      cfg.printMode    = true
		case UPDATE:     cfg.updateMode   = true
		case HELP:       cfg.help         = true
		case VERSION:    cfg.version      = true
		case COLOR:      cfg.color        = true
		case NOMAIN:     cfg.mainScreen = false
		case NOWRITE:    cfg.write      = false
		case EDITOR:     pc.waitForEditor = true
		case DEFALG:     pc.waitForDefAlg = true
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

func executeAlias(cmd *string, pc *parseConfig, cfg *Config) error {
	num := optsFindAlias(cfg.Opts, cmd)
	return executeCommandWithNumber(num, pc, cfg)
}

// For debugging purposes.
func (cfg *Config) Print() {
	fmt.Printf("Review?         : %t\n", cfg.review)
	fmt.Printf("Memorize?       : %t\n", cfg.memorize)
	fmt.Printf("Done?           : %t\n", cfg.done)
	fmt.Printf("Number Enabled? : %t\n", cfg.numberEnabled)
	fmt.Printf("Groups Enabled? : %t\n", cfg.groupsEnabled)
	fmt.Println()
	fmt.Printf("Help    : %t\n", cfg.help)
	fmt.Printf("Version : %t\n", cfg.version)
	fmt.Printf("Color?  : %t\n", cfg.color)
	fmt.Printf("Main Screen?    : %t\n", cfg.mainScreen)
	fmt.Printf("Write?  : %t\n", cfg.write)

	fmt.Printf("Editor  : %s\n", cfg.editor)
	fmt.Printf("Def Alg : %s\n", cfg.defalg)
	fmt.Printf("Number  : %d\n", cfg.number)
	fmt.Printf("Groups  : %v\n", cfg.groups)
	fmt.Printf("Files   : %v\n", cfg.files)
}
