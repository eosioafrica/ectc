package provision_test

import (
	"testing"
	"github.com/eosioafrica/ecte/environment"
	"github.com/eosioafrica/ecte/provision"
	"github.com/eosioafrica/ecte/utils"
)

func TestProvisioner_Provision(t *testing.T) {

	e := utils.CreateTestEnvironment()

	env := environment.New()

	env.Create(e)

	provisioner := provision.New(env)
	provisioner.Provision()

	if provisioner.Err != nil {

		t.Error(provisioner.Err)
	}
}