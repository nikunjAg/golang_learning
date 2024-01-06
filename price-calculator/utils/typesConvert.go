package utils

import (
	"fmt"
	"strconv"
)

func StringToFloat(s string) float64 {
	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println("failed to parse data read from the file", err)
		return 0.0
	}

	return res
}

func StringsToFloatSlice(slice []string) []float64 {

	floatVals := make([]float64, len(slice))
	for i, strPrice := range slice {
		floatPrice := StringToFloat(strPrice)
		floatVals[i] = floatPrice
	}

	return floatVals
}
