package main

import (
	"fmt"
	"math"

	"example.com/investment-calculator/utility"
)

func calculateFutureValue(investmentAmount, expectedReturnRate, numberOfYears, inflation float64) (float64, float64) {
	futureValue := (investmentAmount * math.Pow((1+expectedReturnRate/100), numberOfYears))
	futureRealValue := futureValue / math.Pow(1+inflation/100, numberOfYears)

	return futureValue, futureRealValue
}

func calculateEarningProfitRatio(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit

	return earningsBeforeTax, profit, ratio
}

func profitCalculator() {

	fmt.Println("----------Profit Calculator----------")

	var revenue, expenses, taxRate float64

	revenue, err := utility.CaptureFloatValue("Enter Revenue")
	if err != nil {
		fmt.Println("Err:", err)
		return
	}
	expenses, err = utility.CaptureFloatValue("Enter Expenses")
	if err != nil {
		fmt.Println("Err:", err)
		return
	}
	taxRate, err = utility.CaptureFloatValue("Enter Tax Rate")
	if err != nil {
		fmt.Println("Err:", err)
		return
	}

	earningsBeforeTax, profit, ratio := calculateEarningProfitRatio(revenue, expenses, taxRate)

	// Printf with a format %v -> value for a variable
	fmt.Println("----------------------------------------")
	fmt.Printf("Earnings Before Tax: %.1f\n", earningsBeforeTax)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio (Earnings Before Tax / Profit): %.3f\n", ratio)
}
