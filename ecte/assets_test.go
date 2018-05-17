package ecte_test

import (
	"testing"
	"github.com/eosioafrica/ecte/ecte"
	"os"
)

func TestEnvironment_GetAsset(t *testing.T) {

	en :=  ecte.New()

	dir := en.Config.Dirs.AssetsFull

	t.Log("Attempt to create : %s", dir)

	if err := en.CreateDirectory(dir); err != nil{

		t.Errorf("Test case error : Failed to create file ", err.Error())
		return
	}

	if err := en.DownloadSeedBashInstallAsset(); err != nil{

		t.Errorf("Test case error : Failed to download bash install asset file ", err.Error())
		return
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {

		t.Errorf("Test case error : Could not resolve install directory ", err.Error())
	} else {

		if err := en.DestroyEnvironment(); err !=nil {

			t.Errorf("Test case error : Failed to remove environment ", err.Error())
		}

		t.Log("Test case success : Asset download test case success!")
	}
}