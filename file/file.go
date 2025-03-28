package file

import (
	"encoding/json"
	"fmt"
	"os"
)

type File struct {
	FileName string
}

type PriceData struct {
	Prices []float64 `json:"prices"`
}

func NewFile(fileName string) *File {
	return &File{FileName: fileName}
}

func (f *File) GetPriceFromFile() ([]float64, error) {
	data, err := os.ReadFile(f.FileName)
	if err != nil {
		return []float64{}, err
	}
	var objData PriceData

	err = json.Unmarshal([]byte(data), &objData)

	if err != nil {
		return []float64{}, err
	}

	return objData.Prices, nil
}

func (f *File) WritePriceToFile(prices interface{}, fileName string) {
	pricesText := fmt.Sprintf("%f", prices)
	err := os.WriteFile(fileName, []byte(pricesText), 0644)
	if err != nil {
		fmt.Println("Error", err)
	}
}
