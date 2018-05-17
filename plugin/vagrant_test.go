package plugin_test

import (
	"testing"
	"github.com/eosioafrica/ecte/plugin"
)

func TestVagrant_Provision(t *testing.T) {

	dir := "/home/khosi/go/src/github.com/eosioafrica/ecte/assets"

	vagrant := plugin.Vag{

		Path: dir,
	}

	vagrant.Provision()

	if vagrant.Err != nil {

		t.Errorf("Failed to create vagrant box : ", vagrant.Err)
		return
	}
}