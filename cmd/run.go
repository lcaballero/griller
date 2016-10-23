package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/lcaballero/griller/cmd/task"
	"github.com/lcaballero/griller/config"
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

	conf, err := config.ParseArgs(os.Args[1:])
	if err != nil {
		return
	}

	if conf.Debug && conf.ShowValues {
		b, err := json.MarshalIndent(conf, "", "  ")
		Check(err)
		fmt.Println(string(b))
	}

	Check(task.Generate(conf))
}
