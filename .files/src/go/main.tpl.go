package main

import (
	"fmt"
	"encoding/json"
	"os"
	"{{ .Remote }}/{{ .PackageName }}/cli"
)


func main() {
	conf := cli.ParseArgs(os.Args...)
	bin, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bin))
}

