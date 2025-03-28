package file

import (
	"fmt"
	"os"
	"strconv"
)

type File struct {
	FileName string
}

func NewFile(fileName string) *File {
	return &File{FileName: fileName}
}

func (f *File) GetPriceFromFile() (float64, error) {
	data, err := os.ReadFile(f.FileName)
	if err != nil {
		return 0, err
	}
	dataText := string(data)
	fmt.Println(dataText)

	price, err := strconv.ParseFloat(dataText, 64)
	if err != nil {
		return 0, err
	}

	return price, nil
}
