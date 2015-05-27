package main

import (
	"fmt"
)

var distanceConversionMap = map[string]float64{
	"M":  1,
	"KM": 1000,
	"CM": 0.01,
	"Y":  0.9144,
	"FT": 0.3048,
	"MI": 1609.34,
}

type Distance struct {
	unit  string
	value float64
}

func conversionFactor(fromUnit, toUnit string) float64 {
	return distanceConversionMap[fromUnit] / distanceConversionMap[toUnit]
}

func (d *Distance) Convert(toUnit interface{}) {
	switch val := toUnit.(type) {
	case string:
		d.value = conversionFactor(d.unit, val) * d.value
		d.unit = val
	default:
		fmt.Println("Conversion failed as invalid toUnit data type.")
	}
}
