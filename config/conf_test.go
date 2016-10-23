package config

import (
	. "github.com/lcaballero/exam/assert"
	"testing"
)

func Test_Config_001(t *testing.T) {
	t.Log("Default args should properly default the conf")
	name := "001"
	params := []string{"griller", "project", "--name", name, "--remote", "github.com/saber"}
	conf, err := ParseArgs(params[1:])

	IsNil(t, err)
	IsNotNil(t, conf)
	IsEqStrings(t, conf.Dest, ".")
	IsEqStrings(t, conf.Name, name)
	IsEqStrings(t, conf.Task, "gen")
}
