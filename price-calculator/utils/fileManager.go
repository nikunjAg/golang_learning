package utils

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

func ReadFileByLine(filename string) []string {

	var lines []string
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("reading the file content failed", err)
		return lines
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("reading the file content failed", err)
		// file.Close()
		return lines
	}

	// file.Close()
	return lines
}

func WriteJSONToFile(path string, data any) error {

	file, err := os.Create(path)

	if err != nil {
		return errors.New("unable to write to file")
	}

	defer file.Close()

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		// file.Close()
		return errors.New("failed to convert data to JSON")
	}

	// file.Close()
	return nil
}
