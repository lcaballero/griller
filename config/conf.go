package config

import (
	"github.com/jessevdk/go-flags"
	"os"
	"encoding/json"
)

type Conf struct {
	Dest string `long:"dest" description:"The directory where to output the new project." default:"."`
	Name string `long:"name" description:"Project name, and name of the directory for the newly created go project." required:"true"`
	Task string `long:"task" description:"Task command to run." default:"gen"`
}

func (c *Conf) String() string {
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func ParseArgs(params []string) *Conf {
	conf := &Conf{}
	parser := flags.NewParser(conf, flags.Default)
	_, err := parser.ParseArgs(params)
	if err != nil {
		os.Exit(1)
	}
	return conf
}