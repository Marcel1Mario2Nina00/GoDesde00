package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("ingresa el numero 1 : ")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("el dato ingresado es incorecto " + err.Error())
		}
	}
	fmt.Println("ingresa el numero 2 : ")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("el dato ingresado es incorecto " + err.Error())
		}
	}
	fmt.Println("ingrese la leyenda : ")
	if scanner.Scan() {
		leyenda = scanner.Text()
		fmt.Println(leyenda, numero1*numero2)
	}
}
