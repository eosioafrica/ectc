package plugin_test

import (
	"testing"
	"github.com/eosioafrica/ecte/ecte"
	"github.com/eosioafrica/ecte/plugin"
)

func TestProvisioner_ProvisionHandler(t *testing.T) {

	env := ecte.New()

	provisioner := plugin.New(env)

	if err := provisioner.Provision(); err != nil {

		t.Error(err)
	}
}