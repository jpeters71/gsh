package gsh

import (
	"bozosonparade/gsh"
	"testing"
)

func TestLoadConfigs(t *testing.T) {
	aConfs := gsh.LoadConfigs()
	if len(aConfs) == 0 {
		t.Errorf("LoadConfigs() did not return a value.")
	}

	t.Logf("Configs found:\n")
	for _, conf := range aConfs {
		t.Logf("%s\t%s\n", conf.Name, conf.DefaultSuffix)
	}
}
