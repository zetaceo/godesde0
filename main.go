package main

import (
	"fmt"
	"runtime"

	"github.com/zetaceo/godesde0/ejercicios"
)

func main() {
	/*estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto) */

	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Unix based!", os)
	} else {
		fmt.Println("Un guindos", os)
	}

	numero, texto := ejercicios.Ejercicio("150")

	fmt.Println(numero)
	fmt.Println(texto)
}
