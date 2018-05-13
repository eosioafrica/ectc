package ecte_test

import (
	"testing"
	"github.com/eosioafrica/ecte/ecte"
)



func TestEnvironment_GetDirsToCreate(t *testing.T) {

	en :=  ecte.New()


	dirs := en.GetDirsToCreate();

	if len(dirs) < 1 {

		t.Errorf("Error, no app dirs : ")
		return
	}

	t.Log(dirs)
}
