package Convert

import (
	"RaulOrellanaTestLM/ConvertTypes"
	"fmt"
)

type FactoryImplementation struct {

}

func NewFactoryImplementation() *FactoryImplementation {
	return &FactoryImplementation{}
}

func (f *FactoryImplementation ) Factory(input string, ouput string) (ConvertTypes.IConvertTypes, error){

	if "TEXT" == input{
		if "BINARY" == ouput {
			return ConvertTypes.NewConvertTypeTextToBinaryImplementation(), nil
		}else if "MORSE" == ouput{
			return ConvertTypes.NewConvertTypeTextToMorseImplementation(), nil
		}
	}
	return nil, fmt.Errorf("not valid option")
}
