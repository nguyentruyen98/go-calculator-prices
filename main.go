package main

import (
	"fmt"
	"practice/price-calculator/file"
	"practice/price-calculator/price"
)

func main() {

	file := file.NewFile("prices.json")
	fileData, _ := file.GetPriceFromFile()

	taxList := price.NewTaxList([]float64{0.1, 0.2, 0.3})

	fmt.Println(fileData)

	prices := price.NewPrice(0.1, fileData)
	result := taxList.CalculatorTax(prices.Prices)

	file.WritePriceToFile(result, "outputData.txt")

	fmt.Println(result)
}
