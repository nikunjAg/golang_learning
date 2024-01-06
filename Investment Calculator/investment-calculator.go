package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"example.com/investment-calculator/fileops"
	"example.com/investment-calculator/utility"
)

const INVESTMENT_CALCULATOR_FILE = "investment-result.txt"
const FUTURE_VALUE_LABEL = "Future Value: "
const FUTURE_REAL_VALUE_LABEL = "Future Value (With Inflation): "

func readInvestmentFromFile(filename string) (futureValue, futureRealValue float64, err error) {

	dataBytes, err := fileops.ReadDataFromFile(filename)

	if err != nil {
		return futureValue, futureRealValue, err
	}

	dataTxt := string(dataBytes)
	dataSlices := strings.Split(dataTxt, "\n")

	futureValueTxt := dataSlices[0]
	futureRealValueTxt := dataSlices[1]

	futureValueTxt = strings.Replace(futureValueTxt, FUTURE_VALUE_LABEL, "", 1)
	futureRealValueTxt = strings.Replace(futureRealValueTxt, FUTURE_REAL_VALUE_LABEL, "", 1)

	// Converting Future Value (string) to float64
	futureValue, err = strconv.ParseFloat(futureValueTxt, 64)
	if err != nil {
		return futureValue, futureRealValue, errors.Join(errors.New("unable to convert Future Value to float64"), err)
	}

	// Converting Future Real Value (string) to float64
	futureRealValue, err = strconv.ParseFloat(futureRealValueTxt, 64)
	if err != nil {
		return futureValue, futureRealValue, errors.Join(errors.New("unable to convert Future Real Value to float64"), err)
	}

	return futureValue, futureRealValue, err
}

func writeInvestmentResultToFile(futureValue, futureRealValue float64) error {

	// Convert data to String first
	futureValueText := fmt.Sprintf("%s%.2f\n", FUTURE_VALUE_LABEL, futureValue)
	futureRealValueText := fmt.Sprintf("%s%.2f\n", FUTURE_REAL_VALUE_LABEL, futureRealValue)

	dataText := futureValueText + futureRealValueText

	return fileops.WriteDataToFile(dataText, INVESTMENT_CALCULATOR_FILE)
}

func investmentCalculator() {

	// Prints something to the stdout
	fmt.Println("----------Investment Calculator----------")

	const inflation = 2.3
	var investmentAmount, numberOfYears, expectedReturnRate float64 = 0, 0, 6.5

	investmentAmount, err := utility.CaptureFloatValue("Investement Amount")
	if err != nil {
		fmt.Println("Err:", err)
		return
	}
	expectedReturnRate, err = utility.CaptureFloatValue("Expected Return Rate")
	if err != nil {
		fmt.Println("Err:", err)
		return
	}
	numberOfYears, err = utility.CaptureFloatValue("Number Of Years")
	if err != nil {
		fmt.Println("Err:", err)
		return
	}

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, numberOfYears, inflation)

	fmt.Printf("%s%.2f\n", FUTURE_VALUE_LABEL, futureValue)
	fmt.Printf("%s%.2f\n", FUTURE_REAL_VALUE_LABEL, futureRealValue)

	prevFV, prevFRV, errR := readInvestmentFromFile(INVESTMENT_CALCULATOR_FILE)

	if errR != nil {
		fmt.Println("Err:", errR)
	} else {
		fmt.Printf("Prev %s%.2f\n", FUTURE_VALUE_LABEL, prevFV)
		fmt.Printf("Prev %s%.2f\n", FUTURE_REAL_VALUE_LABEL, prevFRV)
	}

	err = writeInvestmentResultToFile(futureValue, futureRealValue)

	if err == nil {
		fmt.Println("Wrote results successfully to the file.")
	}
}
