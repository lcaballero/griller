package cmd

import (
	"github.com/lcaballero/griller/config"
	"github.com/lcaballero/griller/cmd/task"
	"os"
)

func Check(err error, allow ...error) {
	if err == nil {
		return
	}
	for _, allowed := range allow {
		if allowed == err {
			return
		}
	}
	panic(err)
}

func Run() {
	err := task.NewDotLoader().Load()
	Check(err, task.GrillerDoesNotExistError)

	conf, err := config.ParseArgs(os.Args)
	Check(err)

	switch conf.Task {
	case "gen":
		err = task.Generate(conf)
	default:
		err = task.Generate(conf)
	}

	Check(err)
}
