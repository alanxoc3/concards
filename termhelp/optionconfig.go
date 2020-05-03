package termhelp

import (
   "os"
   "os/user"
   "log"
   "errors"
   "fmt"
   "github.com/alanxoc3/concards/constring"
   "strconv"
)

type Config struct {
	// The various true or false options
	IsReview   bool
	IsMemorize bool
	IsDone     bool
	IsNumber   bool
	IsHelp     bool
	IsVersion  bool
   IsPrint    bool
   IsStream   bool

   Editor     string
   Number     int
   MetaFile   string
	Files      []string
}

func NewConfig() *Config {
   cfg := Config{}

   if val, present = os.LookupEnv("CONCARDS_META"); present {
      cfg.MetaFile = val
   } else if usr, err := user.Current(); err == nil {
      cfg.ConfigFolder = usr.HomeDir + "/.concards-meta"
   } else {
      cfg.MetaFile = ".concards-meta"
   }

   if val, present = os.LookupEnv("EDITOR"); present {
      cfg.Editor = val
   } else {
      cfg.Editor = "vi"
   }

	return &cfg
}

func Help() {
   println(`Usage:
  concards [OPTION]... [FILE|FOLDER]...

Options:
  -r  --review    
  -m  --memorize  
  -d  --done      
  -n  --number #  Limit the number of cards in the program to "#".
  -p  --print     
  -h  --help      
  -E  --editor f  Which editor concards should use. Defaults to "$EDITOR".
  -M  --meta f    Location of concards meta file. Defaults to "$CONCARDS_META" or ~/.concards-meta.

For more details, read the fine man page.
`)
}

func dosomething() {
   // Create new parser object
   parser := argparse.NewParser("concards", "A CLI simple based flashcard parser.")

   // Create flags
   f_review   := parser.Flag("r", "review",   &argparse.Options{Help: "Show cards available to be reviewed."})
   f_memorize := parser.Flag("m", "memorize", &argparse.Options{Help: "Show cards available to be memorized."})
   f_done     := parser.Flag("d", "done",     &argparse.Options{Help: "Show cards not available to be reviewed or memorized."})
   f_print    := parser.Flag("p", "print",    &argparse.Options{Help: "Prints all cards, one line per card."})
   f_help     := parser.Flag("h", "help",     &argparse.Options{Help: "If you need assistance."})

   // Parse input
   err := parser.Parse(os.Args)
   if err != nil {
      // In case of error print error and print usage
      // This can also be done by passing -h or --help flags
      fmt.Print(parser.Usage(err))
      os.Exit(1)
   }
}

// Parses through the given arguments and returns a generated config.
func ParseConfig(args []string) (*Config, error) {
   waitForNum = false
   waitForEditor = false
   waitForMeta = false

   args = args [1:]
	cfg := configInit()
	cfg.Opts = genOptions()
	var tmpArg string // Used in the for loop, for the options passed.

	pc := parseConfig{}

	for _, arg := range args {
      curLen := len(arg)

		if waitForNum {
			waitForNum = false
			num, err := strconv.Atoi(arg)
			if err != nil {
				return nil, errors.New("Error: You didn't pass a number to the number option.")
			}
			cfg.Number = num
		} else if waitForEditor {
			waitForEditor = false
			cfg.Editor = arg
		} else if waitForMeta {
			waitForMeta = false
			cfg.MetaFile = arg
		} else {
         if arg == '-' {

         } else if arg == '--' {

         }
         if cur
         if curLen > 1
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

func CreateDirIfNotExists(dir string) {
   if _, err := os.Stat(dir); !os.IsNotExist(err) {
      return
   }

   if err := os.MkdirAll(dir, 0755); err != nil {
      panic(err)
   }
}

// Set the defaults for the config
func configInit() *Config {
	// Everything besides these are set to false or 0
	var cfg Config
	cfg.Editor = ""

   if usr, err := user.Current(); err != nil {
      log.Fatal( err )
      cfg.ConfigFolder = "."
   } else {
      cfg.ConfigFolder = usr.HomeDir + "/.concards"
   }

   cfg.DatabasePath = cfg.ConfigFolder + "/cards.db"
   cfg.ConfigFile = cfg.ConfigFolder + "/config.yaml"
   CreateDirIfNotExists(cfg.ConfigFolder)

	cfg.Groups = make(map[string]bool)
	return &cfg
}

func (pc *parseConfig) check() bool {
	return pc.waitForGroup || pc.waitForNum || pc.waitForEditor
}
