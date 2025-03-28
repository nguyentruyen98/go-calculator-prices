package price

type TaxIncludedPriceJob struct {
	TaxRate float64
	Prices  []float64
}

func NewPrice(taxRate float64, prices []float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
		Prices:  prices,
	}
}

func (t *TaxIncludedPriceJob) CalculatorTaxRate() []float64 {
	result := make([]float64, len(t.Prices))
	for index, price := range t.Prices {
		result[index] = price * t.TaxRate
	}
	return result
}
