package ecte

import (
	"testing"
)

func TestRun(t *testing.T) {

	e := New()

	e.Run()

	if e.Err != nil {

		t.Errorf("Ecte Test Error : ", e.Err)
	}
}