package main

import (
	"fmt"
)

var moneyExchangeRate = map[string]float64{
	"INR": 0.016,
	"AUD": 0.74,
	"BHD": 2.65,
	"USD": 1,
}

type Money struct {
	currency string
	value    float64
}

type ForexPortfolio []Money

func coversionRate(fromCurrency, toCurrency string) float64 {
	return moneyExchangeRate[fromCurrency] / moneyExchangeRate[toCurrency]
}

func (m *Money) Convert(toCurrency interface{}) {
	switch val := toCurrency.(type) {
	case string:
		m.value = coversionRate(m.currency, val) * m.value
		m.currency = val
	default:
		fmt.Println("Conversion failed as invalid toCurrency data type.")
	}
}

func (ms ForexPortfolio) BulkConvert(toCurrency string) {
	for i := 0; i < len(ms); i++ {
		ms[i].Convert(toCurrency)
	}
	/* if you use range then it created a local copy for each iteration.
	 * Any change performed on the local copy wont relflect in the received object.
	for _, money := range ms {
		money.Convert(toCurrency)
	}
	*/
}
