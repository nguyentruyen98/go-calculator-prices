package price

type TaxList []float64

func NewTaxList(list TaxList) *TaxList {
	return &list
}

func (t *TaxList) CalculatorTax(price []float64) map[float64][]float64 {

	result := make(map[float64][]float64)

	for _, tax := range *t {
		price := NewPrice(tax, price)
		result[tax] = price.CalculatorTaxRate()
	}
	return result
}
