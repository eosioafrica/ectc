package plugin_test

import (
	"testing"
	"github.com/eosioafrica/ecte/plugin"
	"github.com/eosioafrica/ecte/ecte"
)

func TestVagrant_Provision(t *testing.T) {

	dir := ecte.EnvConfig.Dirs.ProvisionersFull

	vagrant := plugin.Vag{

		Path: dir,
	}

	t.Log("The righteous path is : ", vagrant.Path)

	//vagrant.Provision()

	if vagrant.Err != nil {

		t.Errorf("Failed to create vagrant box : ", vagrant.Err)
		return
	}
}