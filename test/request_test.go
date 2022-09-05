package test

import (
	"bark-cli/utils"
	"testing"
)

func TestConfigExistAndHome(t *testing.T) {
	if !utils.ConfigExist(utils.Home()) {
		t.Errorf("utils.Home test failed, $HOME not exist!\n")
	}
}
