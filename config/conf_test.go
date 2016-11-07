package config

import (
	. "github.com/lcaballero/exam/assert"
	"testing"
)

func Test_Config_001(t *testing.T) {
	t.Log("Arguments should come from command line flags")
	params := []string{"template", "--template-name", "templatename", "--project", "projectname"}
	conf, _, err := ParseArgs(params)

	IsNil(t, err)
	IsNotNil(t, conf)
	IsEqStrings(t, conf.Template.Name, "templatename")
	IsEqStrings(t, conf.Template.Project, "projectname")
}
