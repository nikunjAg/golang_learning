package fileops

import (
	"errors"
	"os"
)

const WRITE_PERMISSION = 0644

func ReadDataFromFile(filename string) (string, error) {

	bytesData, err := os.ReadFile(filename)

	if err != nil {
		return "", errors.Join(errors.New("unable to read from the file"), err)
	}

	return string(bytesData), nil
}

func WriteDataToFile(data string, filename string) error {

	err := os.WriteFile(filename, []byte(data), WRITE_PERMISSION)

	if err != nil {
		return errors.Join(errors.New("unable to write to the file"), err)
	}

	return nil
}
