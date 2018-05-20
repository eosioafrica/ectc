package ecte_test

import (
	"testing"
	"github.com/eosioafrica/ecte"
)

func TestRun(t *testing.T) {

	e := ecte.New()

	e.Run()

	if e.Err != nil {

		t.Errorf("Ecte Test Error : ", e.Err)
	}
}