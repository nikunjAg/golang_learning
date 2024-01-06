package taxcalculator

import (
	"fmt"
	"math"

	"example.com/price-calculator/utils"
)

type TaxRate struct {
	Rate        float64
	Prices      []float64
	TaxedPrices map[string]float64
}

func New(rate float64, prices []float64) TaxRate {
	return TaxRate{
		Rate:        rate,
		Prices:      prices,
		TaxedPrices: make(map[string]float64, len(prices)),
	}
}

func (tr *TaxRate) CalculateTax() {

	for _, price := range tr.Prices {
		origPrice := fmt.Sprintf("%.2f", price)
		taxedPrice := price * (1 + tr.Rate)
		tr.TaxedPrices[origPrice] = math.Round((taxedPrice * 100) / 100)
	}
}

func (tr *TaxRate) GetTaxedPricesMap() map[string]float64 {
	return tr.TaxedPrices
}

func (tr *TaxRate) DisplayTaxedPrices() {
	fmt.Printf("%.2f\t:%#v\n", tr.Rate, tr.TaxedPrices)
}

func (tr *TaxRate) SaveToFile() error {
	fileName := fmt.Sprintf("result_%v.json", math.Round(tr.Rate*100))
	return utils.WriteJSONToFile(fileName, tr)
}
