package main

import (
	"fmt"
	"github.com/Marcel1Mario2Nina00/GoDesde00/ejercicios"
	"runtime"
)

func main() {
	//estado, texto := variables.ConviertoATexto(182)
	//fmt.Println(estado)
	//fmt.Println(texto)
	numero, texto := ejercicios.ConvertirString("500")
	fmt.Println(numero)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "Linux" || os == "OS X." {
		fmt.Println("este sistema operativo es Linux , es ", os)
	} else {
		fmt.Println("esto es windows")
	}

	switch os := runtime.GOOS; os {
	case "Linux":
		fmt.Println("esto es linux")
	case "darwin":
		fmt.Println("esto es darwin")
	default:
		fmt.Printf("%s \n", os)
	}

}
