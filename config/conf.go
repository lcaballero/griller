package config

import (
	"github.com/jessevdk/go-flags"
	"encoding/json"
	"os"
)

type Conf struct {
	Dest   string `long:"dest" description:"The directory where to output the new project." env:"GRILLER_DEST"`
	Name   string `long:"name" description:"Project name, and name of the directory for the newly created go project." required:"true"`
	Task   string `long:"task" description:"Task command to run." default:"gen"`
	Remote string `long:"remote" description:"Go repo location.  Example: github.com/lcaballero" required:"true" env:"GRILLER_REMOTE"`
	Debug  bool   `long:"debug" description:"Turns on debug mode which outputs additional information to standard out."`
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
	parser := flags.NewParser(conf, flags.Default)
	_, err := parser.ParseArgs(params)
	if err != nil && len(params) == 1 {
		parser.WriteHelp(os.Stdout)
	}
	if err != nil {
		return nil, err
	}
	return conf, nil
}