package utils

import (
	"os"
	"fmt"
	"io"
	"github.com/mitchellh/go-homedir"
	"path/filepath"
	"io/ioutil"
)

func CreateTestEnvironment() string {

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ecteDir := fmt.Sprintf("%s/%s", home, ".ecte")
	RemoveDirIfExist(ecteDir)
	CreateDirIfNotExist(ecteDir) //Yes I know. Do not ask!

	assetsDir := fmt.Sprintf("%s/%s", home, ".ecte/assets")
	fmt.Println("Setting up environment at : ", ecteDir)

/*
	provsDir := fmt.Sprintf("%s/%s", home, ".ecte/assets/provisioners")
	fmt.Println("Setting up environment at : ", provsDir)
	CreateDirIfNotExist(provsDir)
*/
	eFileSrc := "/home/khosi/go/src/github.com/khosimorafo/assets"
	//eFileDst := fmt.Sprintf("%s/%s", assetsDir, "environment.toml")

	CopyDir(eFileSrc, assetsDir)

	return ecteDir
}

func CreateDirIfNotExist(dir string) {

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func RemoveDirIfExist(dir string) {

	if _, err := os.Stat(dir); os.IsNotExist(err) {

		return
	}
	err := os.RemoveAll(dir)
	if err != nil {
		panic(err)
	}
}

// CopyFile copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file. The file mode will be copied from the source and
// the copied data is synced/flushed to stable storage.
func CopyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	err = out.Sync()
	if err != nil {
		return
	}

	si, err := os.Stat(src)
	if err != nil {
		return
	}
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return
	}

	return
}


// CopyDir recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory must *not* exist.
// Symlinks are ignored and skipped.
func CopyDir(src string, dst string) (err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return
	}
	if err == nil {
		return fmt.Errorf("destination already exists")
	}

	err = os.MkdirAll(dst, si.Mode())
	if err != nil {
		return
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = CopyDir(srcPath, dstPath)
			if err != nil {
				return
			}
		} else {
			// Skip symlinks.
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}

			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return
			}
		}
	}

	return
}