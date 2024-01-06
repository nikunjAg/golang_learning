package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func main() {

	fmt.Println("- Created By", randomdata.Title(randomdata.Male), randomdata.FirstName(randomdata.Male), randomdata.LastName())
	fmt.Println("- Reach Us 24/7", randomdata.PhoneNumber())

	fmt.Println("Welcome to the Calculators World!!")
	for {

		choice := presentChoices()

		switch choice {
		case 1:
			investmentCalculator()
		case 2:
			profitCalculator()
		default:
		}

		if choice == 3 {
			break
		}
	}

	fmt.Println("Good Bye!!")

}
