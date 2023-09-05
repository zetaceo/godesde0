package ejercicios

import (
	"strconv"
)

func Ejercicio(texto string) (int, string) {
	valor, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "Hubo un error: " + err.Error()
	}
	var texto_resultado string

	if valor > 100 {
		texto_resultado = "El valor es mayor a 100"
	} else {
		texto_resultado = "El valor es menor a 100"
	}

	return valor, texto_resultado
}
