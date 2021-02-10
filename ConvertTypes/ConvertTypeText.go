package ConvertTypes

import "fmt"

type ConvertTypeTextToBinaryImplementation struct {
}

func NewConvertTypeTextToBinaryImplementation() *ConvertTypeTextToBinaryImplementation {
	return &ConvertTypeTextToBinaryImplementation{}
}

func (c *ConvertTypeTextToBinaryImplementation) ConvertByType(input string) string {
	var binString string
	for _, c := range input {
		binString += fmt.Sprintf("%.8b", c)
	}

	return binString

}
