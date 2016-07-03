package task

import (
	"testing"
	"github.com/lcaballero/griller/config"
	. "github.com/lcaballero/exam/assert"
)



func Test_Gen_001(t *testing.T) {
	dest := ".test-files"
	args := []string{ "--name", "gen-001", "--dest", dest}
	conf := config.ParseArgs(args)
	err := Generate(conf)
	defer RemoveAll(".", dest)

	IsNil(t, err)
	Exists(t, ".", dest)
	Exists(t, ".", dest + "/gen-001")
}
