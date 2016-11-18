package main

import (
	"fmt"
	"{{ .Remote }}/{{ .PackageName }}/conf"
	"os"
	"encoding/json"
)

func main() {
	fmt.Println("Hello, World!")
	conf := cli.ParseArgs(os.Args[1:]...)
	bin, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bin))
}
