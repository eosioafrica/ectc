package environment_test

import (
	"testing"
	"github.com/eosioafrica/ecte/environment"
)



func TestEnvironment_GetDirsToCreate(t *testing.T) {

	en :=  environment.New()


	dirs := en.GetDirsToCreate();

	if len(dirs) < 1 {

		t.Errorf("Error, no app dirs : ")
		return
	}

	t.Log(dirs)
}
