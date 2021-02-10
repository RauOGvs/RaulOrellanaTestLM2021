package main

import (
	"RaulOrellanaTestLM/Convert"
	"fmt"
)

func main()  {

	factory := Convert.NewFactoryImplementation()
	impl, err :=factory.Factory("TEXT", "MORSE")
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(impl.ConvertByType("RA UL"))


}



