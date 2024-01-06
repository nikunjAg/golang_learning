package main

import (
	"fmt"

	TaxCalculator "example.com/price-calculator/tax-calculator"
	Utils "example.com/price-calculator/utils"
)

func main() {

	taxRates := []float64{0.0, 0.07, 0.10, 0.15}
	strPrices := Utils.ReadFileByLine("prices.txt")
	prices := Utils.StringsToFloatSlice(strPrices)

	fmt.Printf("Prices\t%.2f\n", prices)
	fmt.Println("After Tax Prices\n----------------------------")
	fmt.Println("Tax\t Updated Prices")

	for _, taxRate := range taxRates {
		tR := TaxCalculator.New(taxRate, prices)
		tR.CalculateTax()
		tR.DisplayTaxedPrices()
		err := tR.SaveToFile()
		if err != nil {
			fmt.Println(err)
		}
	}
}
