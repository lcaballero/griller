package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/lcaballero/griller/cmd/task"
	"github.com/lcaballero/griller/config"
	"github.com/lcaballero/griller/embedded"
	"os"
	"strings"
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

	args := os.Args[1:]
	conf, err := config.ParseArgs(args)
	if err != nil {
		return
	}

	if conf.ShowValues {
		b, err := json.MarshalIndent(conf, "", "  ")
		Check(err)
		fmt.Println(string(b))
		return
	}

	if conf.List.IsActive() {
		listTemplateNames()
		return
	}

	err = task.Generate(conf)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func listTemplateNames() {
	set := make(map[string]struct{})
	for _, f := range embedded.AssetNames() {
		n := strings.Index(f, "/")
		template := f[:n]
		set[template] = struct{}{}
	}
	fmt.Println("Available Templates:")
	for k,_ := range set {
		fmt.Printf("  %s\n", k)
	}
}
