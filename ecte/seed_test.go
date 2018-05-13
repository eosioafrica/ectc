package ecte_test

import (
	"testing"
	"github.com/eosioafrica/ecte/ecte"
	"github.com/spf13/viper"
)

func TestEnvironment_SeederHandler(t *testing.T) {

	en :=  ecte.New()

	dir := viper.GetString("directories.assets.full")

	t.Log("Attempt to create : ", dir)
/*
	if err := en.CreateDirectory(dir); err != nil{

		t.Errorf("Test case error : Failed to create file ", err.Error())
		return
	}
*/
	if err := en.SeederHandler(); err != nil{

		t.Errorf("Failed complete seeder process : ", err.Error())
		return
	}

}