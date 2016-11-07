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

func check(err error, allow ...error) {
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

// Run executes the generator causing it to parse the command line
// and generating any boiler plate code based on system setup or
// parsed flags.
func Run() {
	err := task.NewDotLoader().Load()
	check(err, task.ErrGrillerDoesNotExist)

	args := os.Args[1:]
	conf, parser, err := config.ParseArgs(args)
	if err != nil {
		parser.WriteHelp(os.Stdout)
		return
	}

	if conf.ShowValues {
		b, err := json.MarshalIndent(conf, "", "  ")
		check(err)
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
	for k := range set {
		fmt.Printf("  %s\n", k)
	}
}
