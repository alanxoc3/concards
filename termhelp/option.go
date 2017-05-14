package termhelp

import "fmt"

type Option struct {
	Alias       byte
	Command     string
	Param       string
	Description string
}

func newOption(alias byte, command string, param string, description string) *Option {
	o := Option{}
	o.Alias = alias
	o.Command = command
	o.Param = param
	o.Description = description
	return &o
}

func newOptionNoParam(alias byte, command string, description string) *Option {
	return newOption(alias, command, "", description)
}

func (opt *Option) ToString() (ret string) {
	ret += fmt.Sprintf("  ")

	if opt.Alias == 0 {
		ret += fmt.Sprintf("    ")
	} else {
		ret += fmt.Sprintf("-%s, ", string(opt.Alias))
	}

	if opt.Param == "" {
		ret += fmt.Sprintf("--%-10s", opt.Command)
	} else {
		ret += fmt.Sprintf("--%-10s", opt.Command+" "+opt.Param)
	}

	ret += fmt.Sprintf("\t%s\n", opt.Description)

	return ret
}

// Only checks options for alias.
func optsFindAlias(opts []*Option, char byte) int {
	for i, o := range opts {
		if o.Alias == char {
			return i
		}
	}

	return -1
}

// Only checks options for commands.
func optsFindCommand(opts []*Option, str *string) int {
	for i, o := range opts {
		if o.Command == *str {
			return i
		}
	}

	return -1
}
