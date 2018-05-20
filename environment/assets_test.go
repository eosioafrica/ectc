package environment_test

import (
	"testing"
	"github.com/eosioafrica/ecte/environment"
	"os"
)

func TestEnvironment_GetAsset(t *testing.T) {

	en :=  environment.New()

	dir := en.Config.Dirs.AssetsFull

	t.Log("Attempt to create : ", dir)

	en.CreateDirectory(dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {

		t.Errorf("Test case error : Could not resolve install directory ", err.Error())
	}

	/*
	en.DownloadSeedBashInstallAsset()
	if _, err := os.Stat(en.SeedBashScript); os.IsNotExist(err) {

		t.Errorf("Test case error : Could not resolve install directory ", err.Error())
	}
	*/

	en.DestroyEnvironment()

	if en.Err != nil {

		t.Errorf("Test case error : ", en.Err)
	}

	t.Log("Test case success : Asset download test case success!")
}