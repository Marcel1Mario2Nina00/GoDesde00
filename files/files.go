package files

import (
	//	"bufio"
	"bufio"
	"fmt"

	"github.com/Marcel1Mario2Nina00/GoDesde00/ejercicios"

	//"io/ioutil"
	"os"
)

var filename = "./files/txt/tabla.txt"

func GrabarTabla() {
	var texto string = ejercicios.PedirNumero()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("error al crear el archivo", err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}
func SumaTabla() {
	var texto string = ejercicios.PedirNumero()
	if !Append(filename, texto) {
		fmt.Println("erro al concatenar contenido")
	}

}
func Append(file string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("error durante el Appem", err.Error())
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("error durante el WriteString", err.Error())
		return false
	}
	arch.Close()
	return true
}

/*
	func LeoArchivo() {
		archivo, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("Erro al leer el archivo", err.Error())
		}
		fmt.Println(string(archivo))

}
*/
func LeoArchivo() {
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("hubo un error leyendo el archivo", err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> ", registro)
	}
	archivo.Close()
}
