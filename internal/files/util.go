package files

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func Touch(path string) error {
	info, err := os.Stat(path)
	if err == nil {
		// file exists
		if info.IsDir() {
			return fmt.Errorf("%s is a directory", path)
		}
		return nil
	}

	if !errors.Is(err, os.ErrNotExist) {
		// this is an actual error
		return err
	}

	// file does not exist

	// create parent dirs
	err = os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	// create file
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	// set appropriate access bits
	err = file.Chmod(0640)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}
