package main

import (
	"fmt"

	"example.com/investment-calculator/utility"
)

func presentChoices() int {
	fmt.Println("----------Calculators----------")
	fmt.Println("1. Investment Calculator")
	fmt.Println("2. Profit Calculator")
	fmt.Println("3. Exit")
	return utility.CaptureIntValue("Select any option")
}
