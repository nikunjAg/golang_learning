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

	okChannels := make([]chan bool, len(taxRates))
	errChannels := make([]chan error, len(taxRates))

	fmt.Printf("Prices\t%.2f\n", prices)
	fmt.Println("After Tax Prices\n----------------------------")
	fmt.Println("Tax\t Result")

	for index, taxRate := range taxRates {

		okChannels[index] = make(chan bool)
		errChannels[index] = make(chan error)

		tR := TaxCalculator.New(taxRate, prices)
		tR.CalculateTax()
		go tR.SaveToFile(okChannels[index], errChannels[index])
		// tR.DisplayTaxedPrices()
		// if err != nil {
		// 	fmt.Println(err)
		// }
	}

	// for index, okChannel := range okChannels {
	// 	fmt.Println(index)
	// 	<-okChannel
	// }

	// This won't work as in most cases we will success
	// So we will be waiting for data on a channel(errChannel) which is never sent
	// for index, errChannel := range errChannels {
	// 	fmt.Println(index)
	// 	err := <-errChannel
	// }

	for index, taxRate := range taxRates {
		select {
		case err := <-errChannels[index]:
			{
				fmt.Printf("%.2f\t:Failed\n\t %v\n", taxRate, err)
			}
		case <-okChannels[index]:
			fmt.Printf("%.2f\t:Success\n", taxRate)
		}
	}

}
