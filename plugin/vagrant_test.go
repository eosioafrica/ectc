package plugin_test

import (
	"testing"
	"github.com/eosioafrica/ecte/plugin"
	"github.com/eosioafrica/ecte/environment"
)

func TestVagrant_Provision(t *testing.T) {

	en :=  environment.New()
	en.CreateDirectory(en.Config.Dirs.ProvisionersFull)

	vagrant := plugin.Vag{

		Path: environment.EnvConfig.Dirs.ProvisionersFull,
	}

	t.Log("The righteous path is : ", vagrant.Path)

	vagrant.Provision()

	if vagrant.Err != nil {

		t.Errorf("Failed to create vagrant box : ", vagrant.Err)
		return
	}
}