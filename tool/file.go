package tool

import (
	"os"
)

func Directory(filepath string) (err error) {
	_, err = os.Stat(filepath)
	if !os.IsExist(err) {
		err = os.MkdirAll(filepath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
