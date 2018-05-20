package seed_test

import (
	"testing"
	"github.com/eosioafrica/ecte/seed"
)

func TestSeeder_Seed(t *testing.T) {

	seeder := seed.New()

	seeder.ExecDir = "../"

	seeder.Seed()

	if seeder.Err != nil {

		t.Error("Test failure : ", seeder.Err)
	}
}