package gui

import "fmt"

type Option struct {
	Alias       string
	Command     string
	Param       string
	Description string
}

func newOption(alias string, command string, param string, description string) *Option {
	o := Option{}
	o.Alias = alias
	o.Command = command
	o.Param = param
	o.Description = description
	return &o
}

func newOptionNoParam(alias string, command string, description string) *Option {
	return newOption(alias, command, "", description)
}

func optionToString(opt *Option) (ret string) {
	ret += fmt.Sprintf("  ")

	if opt.Alias == "" {
		ret += fmt.Sprintf("    ")
	} else {
		ret += fmt.Sprintf("-%s, ", opt.Alias)
	}

	if opt.Param == "" {
		ret += fmt.Sprintf("--%-10s", opt.Command)
	} else {
		ret += fmt.Sprintf("--%-10s", opt.Command + " " + opt.Param)
	}

	ret += fmt.Sprintf("\t%s\n", opt.Description)

	return ret
}

// True if the option list has either an alias or command.
func optsFindStr(opts []*Option, str *string) int {
	for i,o := range opts {
		if o.Alias == *str || o.Command == *str { return i }
	}

	return -1
}

// Only checks options for alias.
func optsFindAlias(opts []*Option, str *string) int {
	for i,o := range opts {
		if o.Alias == *str { return i }
	}

	return -1
}

// Only checks options for commands.
func optsFindCommand(opts []*Option, str *string) int {
	for i,o := range opts {
		if o.Command == *str { return i }
	}

	return -1
}
