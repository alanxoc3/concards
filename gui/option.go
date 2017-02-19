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
