package ConvertTypes

import (
	"RaulOrellanaTestLM/Convertions"
	"fmt"
	"strings"
)

type ConvertTypeTextToMorseImplementation struct {

}

func NewConvertTypeTextToMorseImplementation() *ConvertTypeTextToMorseImplementation {
	return &ConvertTypeTextToMorseImplementation{}
}


func (c *ConvertTypeTextToMorseImplementation) ConvertByType(input string) string{
	var result string
	inputAux := strings.ToUpper(input)
	var listC []string
	listC=strings.Split(inputAux,"")
	for _, c := range listC {
		result += getMorseByCharacter(c)
	}

	fmt.Println(result)
	return result


}

func getMorseByCharacter(c string) string {
	val:= Convertions.MapMorse()[c]
	return  val


}

