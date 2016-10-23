package task

import (
	. "github.com/lcaballero/exam/assert"
	"testing"
)

func Test_Dotfile_002(t *testing.T) {
	t.Log("Should have .griller file name")
	loader := NewDotLoader()
	loader.Env = func(key string) (string, bool) {
		return "home", true
	}
	file, ok := loader.Filename()

	IsTrue(t, ok)
	IsEqStrings(t, file, "home/.griller")
}

func Test_Dotfile_001(t *testing.T) {
	t.Log("Should have proper default fields")
	loader := NewDotLoader()

	IsNotNil(t, loader.Env)
	IsEqStrings(t, loader.DotName, DefaultGrillerConf)
}
