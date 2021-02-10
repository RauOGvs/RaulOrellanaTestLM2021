package Convert

import "RaulOrellanaTestLM/ConvertTypes"

type IConvertFactory interface {

	Factory(input string, ouput string) (ConvertTypes.IConvertTypes, error)

}
