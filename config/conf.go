package config

import (
	"encoding/json"
	"github.com/jessevdk/go-flags"
	"os"
)

type Conf struct {
	Dest        string   `long:"dest" description:"The directory where to output the new project." env:"GRILLER_DEST"`
	Remote      string   `long:"remote" description:"Go repo location.  Example: github.com/lcaballero" required:"true" env:"GRILLER_REMOTE"`
	Debug       bool     `long:"debug" description:"Turns on debug mode which outputs additional information to standard out."`
	List        Lister   `command:"list" description:"Lists the names of the available templates."`
	Template    Template `command:"template" description:"Generate a boiler plate project from a named template. (See list command)"`
	ShowValues  bool     `long:"show-values" description:"Shows the values parsed from the command line then exits."`
	commandName string   `hidden:"true"`
}

type Lister struct {
	active bool   `hidden:"true"`
	Type   string `long:"type" default:"all"`
}

func (t Lister) IsActive() bool {
	return t.active
}

type Template struct {
	active  bool   `hidden:"true"`
	Name    string `long:"template-name" description:"The name of the template to run." required:"1"`
	Project string `long:"project" description:"Both the project name, and name of the directory for the newly created go project." required:"2"`
}

func (t Template) IsActive() bool {
	return t.active
}

func (c *Conf) String() string {
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func ParseArgs(params []string) (*Conf, error) {
	conf := &Conf{}
	parser := flags.NewParser(conf, flags.HelpFlag)
	_, err := parser.ParseArgs(params)

	// err will be HelpErr for the --help flag
	if err != nil {
		parser.WriteHelp(os.Stdout)
		return nil, err
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
	return conf, nil
}
