package plugin


import (

	"errors"
	"net"
	"strings"
)

// VirtualMachine represents a VBox which can be provisioned using this library.
type VirtualMachine interface {
	GetName() string
	Provision() error
	GetIPs() ([]net.IP, error)
	Destroy() error
	GetState() (string, error)
	Suspend() error
	Resume() error
	Halt() error
	Start() error
	//GetSSH(ssh.Options) (ssh.Client, error)
}

const (
	// VMStarting is the state to use when the VBox is starting
	VMStarting = "starting"
	// VMRunning is the state to use when the VBox is running
	VMRunning = "running"
	// VMHalted is the state to use when the VBox is halted or shutdown
	VMHalted = "halted"
	// VMSuspended is the state to use when the VBox is suspended
	VMSuspended = "suspended"
	// VMPending is the state to use when the VBox is waiting for action to complete
	VMPending = "pending"
	// VMError is the state to use when the VBox is in error state
	VMError = "error"
	// VMUnknown is the state to use when the VBox is unknown state
	VMUnknown = "unknown"
)

var (
	// ErrVMNoIP is returned when a newly provisoned VBox does not get an IP address.
	ErrVMNoIP = errors.New("error getting a new IP for the virtual machine")

	// ErrVMBootTimeout is returned when a timeout occurs waiting for a vm to boot.
	ErrVMBootTimeout = errors.New("timed out waiting for virtual machine")

	// ErrNICAlreadyDisabled is returned when a NIC we are trying to disable is already disabled.
	ErrNICAlreadyDisabled = errors.New("NIC already disabled")

	// ErrFailedToGetNICS is returned when no NICS can be found on the vm
	ErrFailedToGetNICS = errors.New("failed to get interfaces for vm")

	// ErrStartingVM is returned when the VBox cannot be started
	ErrStartingVM = errors.New("error starting VBox")

	// ErrCreatingHD is returned when the VBox cannot be created
	ErrCreatingHD = errors.New("error creating HS")

	// ErrCreatingVM is returned when the VBox cannot be created
	ErrCreatingVM = errors.New("error creating VBox")

	// ErrStartingVM is returned when the VBox cannot be started
	ErrCreatingStorage = errors.New("error creating vm storage")

	// ErrCreatingVM is returned when the VBox cannot be created
	ErrAttachingStorage = errors.New("error attaching vm to storage device")

	// ErrStoppingVM is returned when the VBox cannot be stopped
	ErrStoppingVM = errors.New("error stopping VBox")

	// ErrDeletingVM is returned when the VBox cannot be deleted
	ErrDeletingVM = errors.New("error deleting VBox")

	// ErrVMInfoFailed is returned when the VBox cannot be deleted
	ErrVMInfoFailed = errors.New("error getting information about VBox")

	// ErrVMStateFailed is returned when no state can be parsed for the VBox
	ErrVMStateFailed = errors.New("error getting the state of the VBox")

	// ErrSourceNotSpecified is returned when no source is specified for the VBox
	ErrSourceNotSpecified = errors.New("source not specified")

	// ErrDestNotSpecified is returned when no destination is specified for the VBox
	ErrDestNotSpecified = errors.New("source not specified")

	// ErrSuspendingVM is returned when the VBox cannot be suspended
	ErrSuspendingVM = errors.New("error suspending the VBox")

	// ErrResumingVM is returned when the VBox cannot be resumed
	ErrResumingVM = errors.New("error resuming the VBox")

	// ErrNotImplemented is returned when the operation is not implemented
	ErrNotImplemented = errors.New("operation not implemented")

	// ErrSuspendNotSupported is returned when vm.Suspend() is called, but not supported.
	ErrSuspendNotSupported = errors.New("suspend action not supported")

	// ErrResumeNotSupported is returned when vm.Resume() is called, but not supported.
	ErrResumeNotSupported = errors.New("resume action not supported")

	InfoWaitingForVMSwitchOff = errors.New("waiting for the VBox to switch off...")

	InfoSuccessfulVMSwitchOff = errors.New("VBox has been successfully switched off...")

	InfoAttemptVMDestroy = errors.New("attempting to destroy VBox and delete all artifacts...")

	InfoSuccessfulVMDestroy = errors.New("VBox has been successfully destroyed...")

	// ErrAllocatingMemory is returned when the VBox cannot be created
	ErrAllocatingMemory = errors.New("Failed to allocate memory to machine. ")

	// ErrAssigningMAC is returned when the VBox cannot be created
	ErrAssigningMAC = errors.New("Failed to assign mac address to machine. ")

	// ErrAllocatingMemory is returned when the VBox cannot be created
	ErrAllocatingBootSequence = errors.New("Failed to complete boot sequenc allocation. ")

)

// WrapErrors squashes multiple errors into a single error, separated by ": ".
func WrapErrors(errs ...error) error {

	s := []string{}
	for _, e := range errs {
		if e != nil {
			s = append(s, e.Error())
		}
	}
	return errors.New(strings.Join(s, ": "))
}