package config

import (
	"testing"
	. "github.com/lcaballero/exam/assert"
)

func Test_Config_001(t *testing.T) {
	t.Log("Default args should properly default the conf")
	name := "001"
	params := []string{"griller", "project", "--name", name}
	conf := ParseArgs(params[1:])

	IsNotNil(t, conf)
	IsEqStrings(t, conf.Dest, ".")
	IsEqStrings(t, conf.Name, name)
	IsEqStrings(t, conf.Task, "gen")
}
