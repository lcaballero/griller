package config

import (
	"encoding/json"
	"github.com/jessevdk/go-flags"
	"os"
)

type Conf struct {
	Dest       string   `long:"dest" description:"The directory where to output the new project." env:"GRILLER_DEST"`
	Remote     string   `long:"remote" description:"Go repo location.  Example: github.com/lcaballero" required:"true" env:"GRILLER_REMOTE"`
	Debug      bool     `long:"debug" description:"Turns on debug mode which outputs additional information to standard out."`
	List       bool     `long:"list" description:"Lists the names of the available templates."`
	Template   Template `positional-args:"template" required:"true"`
	ShowValues bool     `long:"show-values" description:"Shows the values parsed from the command line before executing the commend."`
}

type Template struct {
	Name    string `positional-arg-name:"name" description:"The name of the template to run for the project." required:"1"`
	Project string `positional-arg-name:"project" description:"Both the project name, and name of the directory for the newly created go project." required:"2"`
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
	return conf, nil
}
