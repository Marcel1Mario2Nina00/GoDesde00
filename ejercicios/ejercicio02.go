package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PedirNumero() {
	scanner := bufio.NewScanner(os.Stdin)
	var numero int
	var err error
	for {
		fmt.Println("Ingrese un numero: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d ,", numero, i, i*numero)
	}
}
