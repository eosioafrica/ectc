package environment_test

import (
	"testing"
	"github.com/eosioafrica/ecte/utils"
	"github.com/eosioafrica/ecte/environment"
)

func TestEnvironment_Create(t *testing.T) {

	env := environment.New()

	env.DestroyEnvironment()
	if env.Err != nil {

		t.Errorf("Test destroy environment error : ", env.Err.Error())
		return
	}

	e := utils.CreateTestEnvironment()
	env.Create(e)

	if env.Err != nil {

		t.Errorf("Test create environment error : ", env.Err.Error())
		return
	}
}

