package task

import (
	. "github.com/lcaballero/exam/assert"
	"github.com/lcaballero/griller/config"
	"testing"
)

func Test_Gen_003(t *testing.T) {
	t.Log("Should produce data for the template")
	dest := ".test-files"
	args := []string{"--name", "gen003", "--dest", dest, "--remote", "github.com/saber"}
	conf, _ := config.ParseArgs(args)
	data := NewGen(conf).TemplateData()

	IsEqStrings(t, data.PackageName, "gen003")
	IsEqStrings(t, data.Remote, "github.com/saber")
}

func Test_Gen_002(t *testing.T) {
	t.Log("Should create the target directory.")
	dest := ".test-files"
	args := []string{"--name", "gen-001", "--dest", dest, "--remote", "github.com/saber"}
	conf, err := config.ParseArgs(args)
	IsNil(t, err)

	err = Generate(conf)
	defer RemoveAll(".", dest)

	IsNil(t, err)
	Exists(t, dest,
		"gen-001",
		"gen-001/cli/parse.tpl.go",
		"gen-001/conf/config.go",
		"gen-001/.gitignore",
		"gen-001/license",
		"gen-001/main.go",
	)
}

func Test_Gen_001(t *testing.T) {
	t.Log("Should create the target directory.")
	dest := ".test-files"
	args := []string{"--name", "gen-001", "--dest", dest, "--remote", "github.com/saber"}
	conf, err := config.ParseArgs(args)
	IsNil(t, err)

	err = Generate(conf)
	defer RemoveAll(".", dest)

	IsNil(t, err)
	Exists(t, ".", dest)
	Exists(t, ".", dest+"/gen-001")
}
