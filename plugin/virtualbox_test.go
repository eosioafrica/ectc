package plugin_test

import (
	"testing"
	"github.com/eosioafrica/ecte/plugin"
	"fmt"
	"os"
)

var vmName = "TestVM"
var ideMedium = ""
var fileName = ""
var vbox plugin.VBox

// TestMain wraps all tests with the needed initialized
func TestMain(m *testing.M) {

	setup()

	// Run the test suite
	retCode := m.Run()
	os.Exit(retCode)
}


func TestVBox_CreateMachine(t *testing.T) {

	vbox.Destroy()

	vbox.CreateMachine()

	if vbox.Err != nil {

		t.Errorf("Failed to create vbox ", vbox.Err)
		return
	}
}


func TestVBox_CreateStorage(t *testing.T) {

	
	vbox.CreateStorage()

	if vbox.Err != nil {

		t.Errorf("Failed to create machine storage ", vbox.Err)
		return
	}
}

func TestVBox_AllocateMemory(t *testing.T) {

	vbox.AllocateMemory()

	if vbox.Err != nil {

		t.Errorf("Failed to allocate random access memory ", vbox.Err)
		return
	}
}

func TestVBox_BootSequence(t *testing.T) {

	vbox.BootSequence()

	if vbox.Err != nil {

		t.Errorf("Failed to initialize machine boot sequence ", vbox.Err)
		return
	}
}

func TestVBox_CreateInterfaces(t *testing.T) {

	vbox.CreateInterfaces()

	if vbox.Err != nil {

		t.Errorf("Failed to create network interface ", vbox.Err)
		return
	}
}

/*
func TestVBox_Destroy(t *testing.T) {

	vbox := plugin.VBox{

		Name: vmName,
	}

	vbox.Destroy()

	if vbox.Err != nil {

		t.Errorf("Failed to destroy box ", vbox.Err)
		return
	}
}*/

func setup()  {

	dir := "/home/khosi/go/src/github.com/eosioafrica/ecte/assets"
	ideMedium = fmt.Sprintf("%s/%s", dir, "ipxe.iso")
	fileName = fmt.Sprintf("%s/%s/%s.%s", dir, vmName, vmName, "vdi")

	boot := plugin.Boot{

		Boot1: "dvd",
		Boot2: "disk",
		Boot3: "none",
		Boot4: "none",
	}

	nics := []plugin.NIC{
		{

			Idx: 2,
			Mac: "080027AB1001",
			Device: "intnet",
		},
	}

	storage := []plugin.Storage{

		{
			Name: "IDEController",  Bus: "ide",
			Port: 0, Device: 0, Type: "dvddrive", Medium: ideMedium,
		},
		{
			Name: "SATA Controller", Bus: "sata", Controller: "IntelAHCI",
			Port: 0, Device: 0, Type: "hdd", Medium: fileName,
		},
	}

	vbox = plugin.VBox{

		Name: vmName,
		Path: dir,
		OSType: "Linux_64",
		Filename: fileName,
		Config: plugin.MachineConfig{

			NICs: &nics,
			BootSeq: &boot,
			Disk: 10000,
			Memory: 512,
			Storage: &storage,
		},
	}
}