package main

import "os"

func writeToFile(filename string, data string) error {
	return os.WriteFile(filename, []byte(data), 0666)
}

func readFile(filename string) (string, error) {
	var byteArr, err = os.ReadFile(filename)

	return string(byteArr), err
}
