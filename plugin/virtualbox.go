package plugin

import (
	"os/exec"
	"os"
	"bytes"
	"fmt"
	"strings"
	"regexp"
	"time"
)

// VBOXMANAGE is a hardcoded path to VBoxManage to fall back to when it is not
// in the $PATH.
var VBOXMANAGE = "/usr/bin/VBoxManage"

// Regexp for parsing vboxmanage output.
var (
	ipLineRegexp    = regexp.MustCompile(`/VirtualBox/GuestInfo/Net/0/V4/IP`)
	ipAddrRegexp    = regexp.MustCompile(`value: .*, timestamp`)
	timestampRegexp = regexp.MustCompile(`timestamp: \d*`)
	networkRegexp   = regexp.MustCompile(`(?s)Name:.*?VBoxNetworkName`)
	stateRegexp     = regexp.MustCompile(`^State:`)
	runningRegexp   = regexp.MustCompile(`running`)
	backingRegexp   = regexp.MustCompile(`Attachment: NAT`)
	disabledRegexp  = regexp.MustCompile(`disabled$`)
	nicRegexp       = regexp.MustCompile(`^NIC \d\d?:`)
)

// Config represents a config for a VirtualBox VM
type Config struct {

	Disk  		int
	Memory    	int
	NICs 		[]NIC
	Storage		[]Storage
}

// Backing represents a backing for VirtualBox NIC
type Backing int

// NIC represents a Virtualbox NIC
type NIC struct {

	Idx         int			`gorm:"not null;unique"`
	Type       	Backing
	Device 		string
	Mac         string
}

// Represents Storage for the VM
type Storage struct {

	Name		string
	Bus 		string
	Controller  string

	Port		int
	Device		int
	Type		string
	Medium		string
}

// boot sequence for a VirtualBox VM
type Boot struct {

	Boot1  		string 		`gorm:"default:\"none\""`
	Boot2  		string 		`gorm:"default:\"none\""`
	Boot3  		string 		`gorm:"default:\"none\""`
	Boot4  		string 		`gorm:"default:\"none\""`
}

// VM represents a VirtualBox VM
type VM struct {

	Name        string
	Path        string 		// Install path
	ISO         string 		// Source Iso Path
	Config  	Config
	OSType      string
}

// GetName returns the name of the virtual machine
func (vm *VM) GetName() string {

	return vm.Name
}

// Runner is an encapsulation around the vmrun utility.
type Runner interface {

	Run(args ...string) (string, string, error)
	RunCombinedError(args ...string) (string, error)
}

// vboxRunner implements the Runner interface.
type vboxRunner struct {}

var runner Runner = vboxRunner{}

// Will create hard drive and vm, give iso mount points
func (vm *VM) New() error {

	_, err := runner.RunCombinedError("createhd", vm.Name)
	if err != nil {
		// If hd is created succesfully, then create VM
		_, rerr := runner.RunCombinedError("createvm", "--basefolder", vm.Path, "--name",
			vm.Name, "--ostype", vm.OSType, "--register")
		if rerr != nil {
			// If neither succeeds, return both errors.
			return WrapErrors(ErrCreatingVM, err, rerr)
		}
	}
	return nil
}


// Adds and attached storage
func (vm *VM) StorageCtl() error {

	for _, store := range vm.Config.Storage {
		// Create hd storage
		_, err := runner.RunCombinedError("storagectl", vm.Name, "--name", store.Name, "--add", store.Bus)
		if err != nil {

			// If storage is created succesfully, then attach to vm
			_, rerr := runner.RunCombinedError("storageattach", vm.Name, "--storagectl", store.Controller, "--port",
				string(store.Port), "--device", string(store.Device), "--type", store.Type, "--medium", store.Medium)
			if rerr != nil {
				// If neither succeeds, return both errors.
				return WrapErrors(ErrAttachingStorage, err, rerr)
			}
		} else {

			return WrapErrors(ErrAttachingStorage, err)
		}
	}

	return nil
}

func (vm *VM) InterfaceCtl(){


}

// Destroy powers off the VM and deletes its files from disk.
func (vm *VM) Destroy() error {

	err := vm.Halt()
	if err != nil {
		return err
	}

	// wait for vm to be released from lock.

	state, err := vm.GetState()
	i := 0
	// This loop continues while "valid" is true.
	for state != VMHalted {

		fmt.Println(InfoWaitingForVMSwitchOff)
		time.Sleep(2 * time.Second)

		state, _ = vm.GetState()

		if (i >= 5 ){

			return WrapErrors(ErrStoppingVM, err)
		}
		i++
	}

	fmt.Println(InfoSuccessfulVMSwitchOff)

	fmt.Println(InfoAttemptVMDestroy)
	_, err = runner.RunCombinedError("unregistervm", vm.Name, "--delete")
	if err != nil {
		return WrapErrors(ErrDeletingVM, err)
	}

	fmt.Println(InfoSuccessfulVMDestroy)
	return nil
}

// Halt powers off the VM without destroying it
func (vm *VM) Halt() error {

	state, err := vm.GetState()
	if err != nil {
		return err
	}
	if state == VMHalted {
		return nil
	}
	_, err = runner.RunCombinedError("controlvm", vm.Name, "poweroff")
	if err != nil {
		return WrapErrors(ErrStoppingVM, err)
	}
	return nil
}

// Start powers on the VM
func (vm *VM) Start() error {

	_, err := runner.RunCombinedError("startvm", vm.Name)
	if err != nil {
		// If the user has paused the VM it reads as halted but the Start
		// command will fail. Try to resume it as a backup.
		_, rerr := runner.RunCombinedError("controlvm", vm.Name, "resume")
		if rerr != nil {
			// If neither succeeds, return both errors.
			return WrapErrors(ErrStartingVM, err, rerr)
		}
	}
	return nil
}

// GetState gets the power state of the VM being serviced by this driver.
func (vm *VM) GetState() (string, error) {

	stdout, err := runner.RunCombinedError("showvminfo", vm.Name)
	if err != nil {
		return "", WrapErrors(ErrVMInfoFailed, err)
	}
	for _, line := range strings.Split(stdout, "\n") {
		// See if this is a NIC
		if match := stateRegexp.FindStringSubmatch(line); match != nil {
			if match := runningRegexp.FindStringSubmatch(line); match != nil {
				return VMRunning, nil
			}
			return VMHalted, nil
		}
	}
	return VMUnknown, ErrVMStateFailed
}



// Run runs a VBoxManage command.
func (f vboxRunner) Run(args ...string) (string, string, error) {

	var vboxManagePath string
	// If vBoxManage is not found in the system path, fall back to the
	// hard coded path.
	if path, err := exec.LookPath("VBoxManage"); err == nil {
		vboxManagePath = path
	} else {
		vboxManagePath = VBOXMANAGE
	}
	cmd := exec.Command(vboxManagePath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

// RunCombinedError runs a VBoxManage command.  The output is stdout and the the
// combined err/stderr from the command.
func (f vboxRunner) RunCombinedError(args ...string) (string, error) {

	wout, werr, err := f.Run(args...)
	if err != nil {
		if werr != "" {
			return wout, fmt.Errorf("%s: %s", err, werr)
		}
		return wout, err
	}

	return wout, nil
}
