package environment

import (
	"os/exec"
	"os"
	"bytes"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"errors"
	"github.com/eosioafrica/ecte/utils"
)

// Runner is an encapsulation around the vmrun utility.
type Runner interface {

	Run(binary string, args ...string) (string, string, error)
	RunCombinedError(binary string, args ...string) (string, error)
}

type Environment struct{

	SeedBashScript string

	Config 	*EnviroConfig

	Err 	error
}

// envRunner implements the Runner interface.
// envRunner is used to interface to OS commands
type envRunner struct {}

var runner Runner = envRunner{}

/**
*
* Create new environment record.
 */
func New() *Environment {

	return &Environment{

		Config: &EnvConfig,
	}
}

func (env *Environment) Create (sourcePath string) error {

	env.SetSourcePath(sourcePath)

	// Generate new configuration
	env.InitConfig()

	env.CreateAppDirectories()

	//env.CreateUser(env.Config.App.User)

	//env.CreateGroup(env.Config.App.User)

	env.AddUserToGroup(env.Config.App.User, env.Config.App.User)

	env.ValidateEnvironment()

	return env.Err
}

//TODO wrap this inside an interface. Its important that its always implemented.
func (env *Environment) SetSourcePath (path string){

	if _, err := os.Stat(path); os.IsNotExist(err) {

		env.Err = WrapErrors(ErrAppDirDoesNotExist, err)
		return
	}

	env.Config.SourcePath = path
}

func (env *Environment) CreateUser(username string) {

	if env.Err != nil { return }

	var binPath string

	if path, err := exec.LookPath(env.Config.Bash.UserAdd); err != nil {

		env.Err = WrapErrors(ErrOSCommand, err)
	} else {

		binPath = path
	}

	_, err := runner.RunCombinedError(binPath, username)
	if err != nil {

		env.Err = WrapErrors(ErrCreatingUser, err)
	}
}

func (env *Environment) RemoveUser(username string) {

	if env.Err != nil { return }

	var binPath string

	if path, err := exec.LookPath(env.Config.Bash.UserDel); err != nil {

		env.Err = WrapErrors(ErrOSCommand, err)
	} else {

		binPath = path
	}

	_, err := runner.RunCombinedError(binPath,  username) // Delete all files and folders
	if err != nil {

		env.Err = WrapErrors(ErrRemovingUser, err)
	}
}

func (env *Environment) CreateGroup(groupname string) {

	if env.Err != nil { return }

	var binPath string

	if path, err := exec.LookPath(env.Config.Bash.GroupAdd); err != nil {

		env.Err = WrapErrors(ErrOSCommand, err)
	} else {

		binPath = path
	}

	_, err := runner.RunCombinedError(binPath, groupname)
	if err != nil {

		env.Err = WrapErrors(ErrCreatingGroup, err)
	}
}

func (env *Environment) RemoveGroup(groupname string)  {

	if env.Err != nil { return }

	var binPath string

	if path, err := exec.LookPath(env.Config.Bash.GroupDel); err != nil {

		env.Err = WrapErrors(ErrOSCommand, err)
	} else {

		binPath = path
	}

	_, err := runner.RunCombinedError(binPath, "-r", groupname) // Delete all files and folders
	if err != nil {

		env.Err = WrapErrors(ErrRemovingUser, err)
	}
}

func (env *Environment) AddUserToGroup(username string, group string) {

	var binPath string

	if path, err := exec.LookPath(env.Config.Bash.UserMod); err != nil {

		env.Err = WrapErrors(ErrOSCommand, err)
	} else {

		binPath = path
	}

	_, err := runner.RunCombinedError(binPath,"-aG", group, username)
	if err != nil {

		env.Err = WrapErrors(ErrAddingUserToGroup, err)
	}

}

func (env *Environment) AddUserToSudoers(username string) {

	var binPath string

	if path, err := exec.LookPath(env.Config.Bash.Echo); err != nil {

		env.Err = WrapErrors(ErrOSCommand, err)
	} else {

		binPath = path
	}

	_, err := runner.RunCombinedError(binPath,"-aG", username)
	if err != nil {

		env.Err = WrapErrors(ErrAddingUserToSudoers, err)
	}
}

func (env *Environment) CreateAppDirectories()  {

	if env.Err != nil { return }

	dirs := env.GetDirsToCreate()
	if len(dirs) < 1 {

		env.Err = WrapErrors(ErrAppDirsConfigNoDirectories)
	}

	for i := range dirs {

		env.CreateDirectory(dirs[i])
	}
}

func (env *Environment) CreateDirectory(dir string) {

	if env.Err != nil { return }

	utils.CreateDirIfNotExist(dir)
}

func (env *Environment) RemoveDirectory(dir string) {

	if env.Err != nil { return }

	utils.RemoveDirIfExist(dir)
}

func (env *Environment) DestroyEnvironment() {

	//env.RemoveUser("ecte")

	//env.RemoveGroup("ecte")

	env.RemoveAppFolder()
}

func (env *Environment) RemoveAppFolder() {


	if env.Err != nil { return }

	env.RemoveDirectory(env.Config.Dirs.AppPathFull)
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

// An explicit check to verify that the environment has all the required folders.
func (env *Environment) ValidateEnvironment(){

	if env.Err != nil { return }

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if home != env.Config.Home {

		message := fmt.Sprintf("App home is no the same as user home. Verify seeder execution.")
		env.Err = WrapErrors(errors.New(message))
		return
	}

	if _, err := os.Stat(fmt.Sprintf("%s/%s", home, ".ecte")); os.IsNotExist(err) {

		message := fmt.Sprintf("%s does not exist. Verify seeder execution.", fmt.Sprintf("%s/%s", home, ".ecte/bin"))
		env.Err = WrapErrors(errors.New(message))
		return
	}

	if _, err := os.Stat(fmt.Sprintf("%s/%s", home, ".ecte/assets/provisioners")); os.IsNotExist(err) {

		message := fmt.Sprintf("%s does not exist.  Verify seeder execution.", fmt.Sprintf("%s/%s", home, ".ecte/assets/provisioners"))
		env.Err = WrapErrors(errors.New(message))
		return
	}

	if _, err := os.Stat(fmt.Sprintf("%s/%s", home, ".ecte/virtualbox")); os.IsNotExist(err) {

		message := fmt.Sprintf("%s does not exist.", fmt.Sprintf("%s/%s", home, ".ecte/virtualbox"))
		env.Err = WrapErrors(errors.New(message))
		return
	}

	if _, err := os.Stat(fmt.Sprintf("%s/%s", home, ".ecte/bin")); os.IsNotExist(err) {

		message := fmt.Sprintf("%s does not exist.  Verify seeder execution.", fmt.Sprintf("%s/%s", home, ".ecte/bin"))
		env.Err = WrapErrors(errors.New(message))
		return
	}
}
