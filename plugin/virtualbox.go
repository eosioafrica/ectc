package plugin

import (
	"os/exec"
	"os"
	"bytes"
	"fmt"
	"strings"
	"regexp"
	"time"
	"errors"
	"strconv"
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

// MachineConfig represents a config for a VirtualBox VBox
type MachineConfig struct {

	Disk  		int
	Memory    	int
	NICs 		*[]NIC
	Storage		*[]Storage
	BootSeq     *Boot
}

// Backing represents a backing for VirtualBox NIC
type Backing int

// NIC represents a Virtualbox NIC
type NIC struct {

	Idx         int
	Type       	Backing
	Device 		string
	Mac         string
}

// Represents Storage for the VBox
type Storage struct {

	Name		string
	Bus 		string
	Controller  string

	Port		int
	Device		int
	Type		string
	Medium		string
}

// boot sequence for a VirtualBox VBox
type Boot struct {

	Boot1  		string
	Boot2  		string
	Boot3  		string
	Boot4  		string
}

// VBox represents a VirtualBox machine
type VBox struct {

	Name   string
	Path   string 		// Install path
	ISO    string 		// Source Iso Path
	OSType string
	Filename string
	Config MachineConfig

	Err 		error
}

// GetName returns the name of the virtual machine
func (vbox *VBox) GetName() string {

	return vbox.Name
}

// Runner is an encapsulation around the vmrun utility.
type Runner interface {

	Run(args ...string) (string, string, error)
	RunCombinedError(args ...string) (string, error)
}

// vboxRunner implements the Runner interface.
type vboxRunner struct {}

var runner Runner = vboxRunner{}

func (vbox *VBox) Create () error {

	vbox.CreateMachine()
	vbox.CreateStorage()
	vbox.AllocateMemory()
	vbox.BootSequence()
	vbox.CreateInterfaces()

	return vbox.Err
}

// Will create hard drive and vm, give iso mount points
func (vbox *VBox) CreateMachine () {

	if vbox.Err != nil { return }

	_, err := runner.RunCombinedError("createhd", "--filename", vbox.Filename, "--size", "20000")
	if err != nil {

		vbox.Err = WrapErrors(ErrCreatingHD, err)

	} else {
		// If hd is created succesfully, then create VBox
		_, rerr := runner.RunCombinedError("createvm", "--basefolder", vbox.Path, "--name",
			vbox.Name, "--ostype", vbox.OSType, "--register")
		if rerr != nil {
			// If neither succeeds, return both errors.
			vbox.Err = WrapErrors(ErrCreatingVM, err, rerr)
		}
	}
}

// Adds and attached storage
func (vbox *VBox) CreateStorage() {

	if vbox.Err != nil { return }

	strg := vbox.Config.Storage

	for _, store := range *strg {
		// Create hd storage
		_, err := runner.RunCombinedError("storagectl", vbox.Name, "--name", store.Name, "--add", store.Bus)
		if err != nil {

			message := fmt.Sprintf("IDE at %s" , store.Name)
			vbox.Err =  WrapErrors(ErrAttachingStorage,errors.New(message), err)

		} else {

			// If storage is created succesfully, then attach to vbox
			_, rerr := runner.RunCombinedError("storageattach", vbox.Name, "--storagectl", store.Name, "--port",
				"0", "--device", "0", "--type", store.Type, "--medium", store.Medium)
			if rerr != nil {
				// If neither succeeds, return both errors.
				message := fmt.Sprintf("SATA at %s" , store.Name)
				vbox.Err = WrapErrors(ErrAttachingStorage, errors.New(message), err, rerr)
			}
		}
	}

}

func (vbox *VBox) CreateInterfaces(){

	if vbox.Err != nil { return }

	nics := vbox.Config.NICs

	for _, nic := range *nics {

		// Virtualbox indexes start at 1
		nicLabel := fmt.Sprintf("--nic%d", nic.Idx)
		macLabel := fmt.Sprintf("--macaddress%d", nic.Idx)

		_, err := runner.RunCombinedError("modifyvm", vbox.Name, nicLabel, nic.Device)
		if err != nil {

			vbox.Err =  WrapErrors(ErrAttachingStorage, err)
		}

		_, err = runner.RunCombinedError("modifyvm", vbox.Name, macLabel, nic.Mac)
		if err != nil {

			vbox.Err =  WrapErrors(ErrAttachingStorage, err)
		}
	}
}

func (vbox *VBox) AllocateMemory(){

	if vbox.Err != nil { return }

	_, err := runner.RunCombinedError("modifyvm", vbox.Name, "--memory",
			strconv.Itoa(vbox.Config.Memory), "--vram", "128")
	if err != nil {

		vbox.Err =  WrapErrors(ErrAllocatingMemory, err)
	}
}

func (vbox *VBox) BootSequence()  {

	if vbox.Err != nil { return }

	_, err := runner.RunCombinedError("modifyvm", vbox.Name,
		"--boot1", vbox.Config.BootSeq.Boot1,
		"--boot2", vbox.Config.BootSeq.Boot2,
		"--boot3", vbox.Config.BootSeq.Boot3,
		"--boot4", vbox.Config.BootSeq.Boot4 )

	if err != nil {

		vbox.Err =  WrapErrors(ErrAllocatingBootSequence, err)
	}
}

// Destroy powers off the VBox and deletes its files from disk.
func (vbox *VBox) Destroy() error {

	err := vbox.Halt()
	if err != nil {
		return err
	}

	// wait for vbox to be released from lock.

	state, err := vbox.GetState()
	i := 0
	// This loop continues while "valid" is true.
	for state != VMHalted {

		fmt.Println(InfoWaitingForVMSwitchOff)
		time.Sleep(2 * time.Second)

		state, _ = vbox.GetState()

		if (i >= 5 ){

			return WrapErrors(ErrStoppingVM, err)
		}
		i++
	}

	fmt.Println(InfoSuccessfulVMSwitchOff)

	fmt.Println(InfoAttemptVMDestroy)
	_, err = runner.RunCombinedError("unregistervm", vbox.Name, "--delete")
	if err != nil {
		return WrapErrors(ErrDeletingVM, err)
	}

	fmt.Println(InfoSuccessfulVMDestroy)
	return nil
}

// Halt powers off the VBox without destroying it
func (vbox *VBox) Halt() error {

	state, err := vbox.GetState()
	if err != nil {
		return err
	}
	if state == VMHalted {
		return nil
	}
	_, err = runner.RunCombinedError("controlvm", vbox.Name, "poweroff")
	if err != nil {
		return WrapErrors(ErrStoppingVM, err)
	}
	return nil
}

// Start powers on the VBox
func (vbox *VBox) Start() error {

	_, err := runner.RunCombinedError("startvm", vbox.Name)
	if err != nil {
		// If the user has paused the VBox it reads as halted but the Start
		// command will fail. Try to resume it as a backup.
		_, rerr := runner.RunCombinedError("controlvm", vbox.Name, "resume")
		if rerr != nil {
			// If neither succeeds, return both errors.
			return WrapErrors(ErrStartingVM, err, rerr)
		}
	}
	return nil
}

// GetState gets the power state of the VBox being serviced by this driver.
func (vbox *VBox) GetState() (string, error) {

	stdout, err := runner.RunCombinedError("showvminfo", vbox.Name)
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
// combined Err/stderr from the command.
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
