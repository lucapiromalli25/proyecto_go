package ejercicios

import (
	"strconv"
)

func Ejercicio01(numero string) (int, string) {
	n, err := strconv.Atoi(numero)
	if err != nil {
		return 0, "Hubo un error inesperado" + err.Error()
	}
	if n > 100 {
		return n, "El numero es mayor a 100"
	} else {
		return n, "El numero es menor o igual 100"
	}
}
