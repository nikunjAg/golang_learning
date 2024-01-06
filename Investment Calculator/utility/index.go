package utility

import (
	"errors"
	"fmt"
)

func CaptureIntValue(str string) (val int) {
	fmt.Print(str, ": ")
	fmt.Scan(&val)
	return val
}

func CaptureFloatValue(str string) (val float64, err error) {
	fmt.Print(str, ": ")
	fmt.Scan(&val)

	if val <= 0 {
		return val, errors.New("please enter a positive value")
	}

	return val, err
}
