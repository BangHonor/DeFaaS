package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CheckDir ...
func CheckDir(dirPath string) error {
	stat, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("directory '%s' does not exist", dirPath)
	}
	if !stat.IsDir() {
		return fmt.Errorf("'%s' should be a directory", dirPath)
	}
	return nil
}

// IsDirEmpty ...
// https://stackoverflow.com/questions/30697324/how-to-check-if-directory-on-path-is-empty
func IsDirEmpty(dirPath string) (bool, error) {

	if err := CheckDir(dirPath); err != nil {
		return false, err
	}

	dir, err := os.Open(dirPath)
	if err != nil {
		return false, err
	}
	defer dir.Close()

	_, err = dir.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}

// ClearDir ...
// https://stackoverflow.com/questions/33450980/how-to-remove-all-contents-of-a-directory-using-golang
func ClearDir(dirPath string) error {

	if err := CheckDir(dirPath); err != nil {
		return err
	}

	dir, err := os.Open(dirPath)
	if err != nil {
		return err
	}
	defer dir.Close()

	names, err := dir.Readdirnames(-1)
	if err != nil {
		return err
	}

	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dirPath, name))
		if err != nil {
			return err
		}
	}

	return nil
}
