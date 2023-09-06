package ejercicios

import (
	"strconv"
)

func ConvertirString(numeroS string) (int, string) {
	num, error := strconv.Atoi(numeroS)
	if error != nil {
		return 0, "hubo un error" + error.Error()
	}
	if num > 100 {
		return num, "es mayor a 100"
	} else {
		return num, "es menor a 100"
	}
}
