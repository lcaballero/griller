package cmd

import (
	"github.com/lcaballero/griller/config"
	"github.com/lcaballero/griller/cmd/task"
	"os"
)

func Run() {
	conf := config.ParseArgs(os.Args)
	switch conf.Task {
	case "gen":
		task.Generate(conf)
	default:
		task.Generate(conf)
	}
}




