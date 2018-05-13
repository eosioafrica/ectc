package ecte

import (
	"os/exec"
	"os"
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

// Runner is an encapsulation around the vmrun utility.
type Runner interface {

	Run(binary string, args ...string) (string, string, error)
	RunCombinedError(binary string, args ...string) (string, error)
}

type Environment struct{


	SeedBashScript string


}

// envRunner implements the Runner interface.
type envRunner struct {}

var runner Runner = envRunner{}

/**
*
* Create new environment record.
 */
func New() *Environment {

	return &Environment{}
}

func (env *Environment) CreateUser(username string) error {

	var binPath string

	if path, err := exec.LookPath(UserAdd()); err != nil {

		return WrapErrors(ErrOSCommand, err)
	} else {

		binPath = path
	}

	_, err1 := runner.RunCombinedError(binPath, username)
	if err1 != nil {

		return WrapErrors(ErrCreatingUser, err1)
	}

	return nil
}

func (env *Environment) RemoveUser(username string) error {

	var binPath string

	if path, err := exec.LookPath(UserDel()); err != nil {

		return WrapErrors(ErrOSCommand, err)
	} else {

		binPath = path
	}

	_, err1 := runner.RunCombinedError(binPath,  username) // Delete all files and folders
	if err1 != nil {

		return WrapErrors(ErrRemovingUser, err1)
	}

	return nil
}

func (env *Environment) CreateGroup(groupname string) error {

	var binPath string

	if path, err := exec.LookPath(GroupAdd()); err != nil {

		return WrapErrors(ErrOSCommand, err)
	} else {

		binPath = path
	}

	_, err1 := runner.RunCombinedError(binPath, groupname)
	if err1 != nil {

		return WrapErrors(ErrCreatingGroup, err1)
	}

	return nil
}

func (env *Environment) RemoveGroup(groupname string) error {

	var binPath string

	if path, err := exec.LookPath(GroupDel()); err != nil {

		return WrapErrors(ErrOSCommand, err)
	} else {

		binPath = path
	}

	_, err1 := runner.RunCombinedError(binPath, "-r", groupname) // Delete all files and folders
	if err1 != nil {

		return WrapErrors(ErrRemovingUser, err1)
	}

	return nil
}

func (env *Environment) AddUserToGroup(username string, group string) error {

	var binPath string

	if path, err := exec.LookPath(UserMod()); err != nil {

		return WrapErrors(ErrOSCommand, err)
	} else {

		binPath = path
	}

	_, err1 := runner.RunCombinedError(binPath,"-aG", group, username)
	if err1 != nil {

		return WrapErrors(ErrAddingUserToGroup, err1)
	}

	return nil
}

func (env *Environment) AddUserToSudoers(username string) error {

	var binPath string

	if path, err := exec.LookPath(Echo()); err != nil {

		return WrapErrors(ErrOSCommand, err)
	} else {

		binPath = path
	}

	_, err1 := runner.RunCombinedError(binPath,"-aG", username)
	if err1 != nil {

		return WrapErrors(ErrAddingUserToSudoers, err1)
	}
	return nil
}

func (env *Environment) CreateAppDirectories() error {

	dirs := env.GetDirsToCreate()
	if len(dirs) < 1 {

		return WrapErrors(ErrAppDirsConfigNoDirectories)
	}

	for i := range dirs {

		if err := env.CreateDirectory(dirs[i]); err != nil {

			return WrapErrors(ErrCreatingAppDirectories, err)
		}
	}
	return nil
}

func (env *Environment) CreateDirectory(dir string) error {

	var binPath string

	if path, err := exec.LookPath(MkDir()); err != nil {

		return WrapErrors(ErrOSCommand, err)
	} else {

		binPath = path
	}

	_, err1 := runner.RunCombinedError(binPath, "-p", dir)
	if err1 != nil {

		return WrapErrors(ErrCreatingDirectory, err1)
	}
	return nil
}

func (env *Environment) RemoveDirectory(dir string) error {

	var binPath string

	if path, err := exec.LookPath(RM()); err != nil {

		return WrapErrors(ErrOSCommand, err)
	} else {

		binPath = path
	}

	_, err1 := runner.RunCombinedError(binPath, "-r", dir)
	if err1 != nil {

		return WrapErrors(ErrRemovingDirectory, err1)
	}
	return nil
}

func (env *Environment) DestroyEnvironment() error {

	var binPath string

	if path, err := exec.LookPath(RM()); err != nil {

		return WrapErrors(ErrOSCommand, err)
	} else {

		binPath = path
	}


	_, err1 := runner.RunCombinedError(binPath, "-r", DirContent())
	if err1 != nil {

		return WrapErrors(ErrDeletingEnvironment, err1)
	}
	return nil
}

func (f envRunner) RunCombinedError(binary string, args ...string) (string, error) {

	wout, werr, err := f.Run(binary, args...)
	if err != nil {
		if werr != "" {
			return wout, fmt.Errorf("%s: %s", err, werr)
		}
		return wout, err
	}
	return wout, nil
}

// Run shell command.
func (f envRunner) Run(binary string, args ...string) (string, string, error) {

	cmd := exec.Command(binary, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func (env envRunner) checkIfSudoUser(){}

func MkDir () string {

	return viper.GetString("CMDMkDir")
}

func UserAdd () string {

	return viper.GetString("CMDUserAdd")
}

func UserDel () string {

	return viper.GetString("CMDUserDel")
}

func GroupAdd () string {

	return viper.GetString("CMDGroupAdd")
}

func GroupDel () string {

	return viper.GetString("CMDGroupDel")
}

func UserMod () string {

	return viper.GetString("CMDUserMod")
}

func Echo () string {

	return viper.GetString("CMDEcho")
}

func RM () string {

	return viper.GetString("CMDRM")
}

// Directories

func DirContent () string {

	if !viper.IsSet("directories.content") {
		log.Fatal("Missing content directory location")
	}
	return viper.GetString("directories.content")
}

func DirVirtualBox () string {

	if !viper.IsSet("directories.virtualbox") {
		log.Fatal("Missing virtualbox directory location")
	}
	return viper.GetString("directories.virtualbox")
}

func DirAssets () string {

	if !viper.IsSet("directories.assets") {
		log.Fatal("Missing assets directory location")
	}
	return viper.GetString("directories.assets")
}

func DirBin () string {

	if !viper.IsSet("directories.bin") {
		log.Fatal("Missing bin directory location")
	}
	return viper.GetString("directories.bin")
}