package ecte

import (

	"errors"
	"strings"
)

var (
	// ErrCreatingUser
	ErrCreatingUser = errors.New("Error while creating user. Please ensure that you have root privilegdes on your system.")

	// ErrRemovingUser
	ErrRemovingUser = errors.New("Failed to remove user from the system.")

	// ErrCreatingUser
	ErrCreatingGroup = errors.New("Error while creating group. Please ensure that you have root privilegdes on your system.")

	// ErrRemovingUser
	ErrRemovingGroup = errors.New("Failed to remove group from the system.")

	// ErrAddingUserToGroup
	ErrAddingUserToGroup = errors.New("Failed to add user to group.")

	// ErrAddingUserToSudoers
	ErrAddingUserToSudoers = errors.New("Failed to add user to sudoers.")

	// InfoUserCreated
	InfoUserCreated = errors.New("User has been created.")

	// ErrCreatingDirectory
	ErrCreatingDirectory = errors.New("Error while creating directory %s. ")

	// ErrCreatingAppDirectories
	ErrCreatingAppDirectories = errors.New("Error while creating app directories.")

	// ErrAppDirsConfigNoDirectories
	ErrAppDirsConfigNoDirectories = errors.New("No app directories. Please check your config. ")

	// InfoDirectoryCreated
	InfoDirectoryCreated = errors.New("Directory has been created.")

	// InfoEnvironmentBaseDirectory
	InfoEnvironmentBaseDirectory = errors.New("Will create new ecte environment at ")

	// ErrOSCommand i
	ErrOSCommand = errors.New("OS command could not be resolved.  ")

	// ErrRemovingDirectory
	ErrRemovingDirectory = errors.New("Failed to delete directory.")

	// ErrDeletingEnvironment
	ErrDeletingEnvironment = errors.New("Failed to delete content directory. ")

	// ErrInvalidURLSubmitted
	ErrInvalidURLSubmitted = errors.New("Invalid URL submitted. ")

	// ErrDownloadingBashSeedAsset
	ErrDownloadingBashSeedAsset = errors.New("Failed to download bash script seed assets. ")

	// InfoDownloadingBashSeedAssetSuccess
	InfoDownloadingBashSeedAssetSuccess = errors.New("Successfully downloaded bash script seed assets. ")

	// ErrRunBashDependencyInstallation
	ErrAssetDirDoesNotExist = errors.New("Failed to resolve asset directory. ")

	// InfoRunBashDependencyInstallationSuccess
	InfoRunBashDependencyInstallationSuccess = errors.New("Successfully executed dependency bash script. ")

	// ErrRunBashDependencyInstallation
	ErrRunBashDependencyInstallation = errors.New("Failed to execute dependency bash script ")
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
