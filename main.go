package main

import (
	"fmt"

	"github.com/Marcel1Mario2Nina00/GoDesde00/variables"
)

func main() {
	estado, texto := variables.ConviertoATexto(182)
	fmt.Println(estado)
	fmt.Println(texto)
}
