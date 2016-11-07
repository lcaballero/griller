package config

import (
	"encoding/json"
	"github.com/jessevdk/go-flags"
)

// Conf represents values exported from the command line.
type Conf struct {
	Dest        string   `long:"dest" description:"The directory where to output the new project." env:"GRILLER_DEST"`
	Remote      string   `long:"remote" description:"Go repo location.  Example: github.com/lcaballero" required:"true" env:"GRILLER_REMOTE"`
	Debug       bool     `long:"debug" description:"Turns on debug mode which outputs additional information to standard out."`
	List        Lister   `command:"list" description:"Lists the names of the available templates."`
	Template    Template `command:"template" description:"Generate a boiler plate project from a named template. (See list command)"`
	ShowValues  bool     `long:"show-values" description:"Shows the values parsed from the command line then exits."`
	commandName string   `hidden:"true"`
}

// Lister represents the List command from the cli.
type Lister struct {
	active bool   `hidden:"true"`
	Type   string `long:"type" default:"all"`
}

// IsActive indicates (when true) that this command is the command issued
// at the command line rather than any of the other values.
func (t Lister) IsActive() bool {
	return t.active
}

// Template represets the template command from the cli
type Template struct {
	active  bool   `hidden:"true"`
	Name    string `long:"template-name" description:"The name of the template to run." required:"1"`
	Project string `long:"project" description:"Both the project name, and name of the directory for the newly created go project." required:"2"`
}

// IsActive indicates that a template command was issued at the
// command line.
func (t Template) IsActive() bool {
	return t.active
}

// String provides and easy way to visualize the conf structure as it was
// parsed from the command line.
func (c *Conf) String() string {
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(b)
}

// ParseArgs takes in the command line tokens (minus the first) and returns
// a new Conf based on the values parsed, a parser used to parse those values
// and an error if one should occur during the parsing phase.  If an error
// does occur then the parse provided can be used to show the help/usage.
func ParseArgs(params []string) (*Conf, *flags.Parser, error) {
	conf := &Conf{}
	parser := flags.NewParser(conf, flags.HelpFlag)
	_, err := parser.ParseArgs(params)

	// err will be HelpErr for the --help flag
	if err != nil {
		return nil, parser, err
	}
	if len(params) > 0 {
		cmd := params[0]
		switch cmd {
		case "list":
			conf.commandName = cmd
			conf.List.active = true
		case "template":
			conf.commandName = cmd
			conf.Template.active = true
		}
	}
	return conf, parser, nil
}
