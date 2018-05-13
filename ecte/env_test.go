package ecte_test

import (
	"testing"
	"os"
	"github.com/eosioafrica/ecte/ecte"
	"os/user"
)

var baseDir = ".ecte"


func TestEnvironment_CreateUser(t *testing.T) {

	en :=  ecte.New()

	username := "ectc"

	if err := en.CreateUser(username); err != nil{

		t.Errorf(err.Error())
		return
	}

	if _, err := user.Lookup(username); err != nil {
		t.Errorf("Failed to create user. ", err.Error())
		return
	} else {

		t.Log("User creation success!")
		if err := en.RemoveUser(username); err !=nil {

			t.Errorf("Failed to remove user ", err.Error())
			return
		}
	}
}

func TestEnvironment_AddUserToSudoers(t *testing.T) {

}

func TestEnvironment_CreateDirectory(t *testing.T) {

	en :=  ecte.New()

	dir := baseDir

	if err := en.CreateDirectory(dir); err != nil{

		t.Errorf("Failed to create file ", err.Error())
		return
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		t.Errorf("Failed to create file ", err.Error())
	} else {

		t.Log("File creation success!")
		if err := en.RemoveDirectory(dir); err !=nil {

			t.Errorf("Failed to remove directory ", err.Error())
		}
	}
}

func TestEnvironment_CreateAppDirectories(t *testing.T) {

	en :=  ecte.New()

	if err := en.CreateAppDirectories(); err != nil{

		t.Errorf("Failed to create app dirs : ", err.Error())
		return
	}
}

func TestEnvironment_DestroyEnvironment(t *testing.T) {

	en :=  ecte.New()

	dir := baseDir

	if err := en.CreateDirectory(dir); err != nil{

		t.Errorf("Failed to create file ", err.Error())
		return
	}

	if err := en.DestroyEnvironment(); err !=nil {

		t.Errorf("Failed to destroy environment ", err.Error())
	}
}