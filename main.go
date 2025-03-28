package main

import (
	"fmt"
	"practice/price-calculator/file"
	"practice/price-calculator/price"
)

func main() {

	file := file.NewFile("prices.txt")

	fileData, _ := file.GetPriceFromFile()

	prices := price.NewPrice(0.1, []float64{123, 123, 123})
	result := prices.CalculatorTaxRates()
	fmt.Println(result)
	fmt.Println(fileData)
}
