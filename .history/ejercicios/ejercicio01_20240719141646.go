package ejercicios

import (
	"strconv"
)

func Ejercicio01(numero string) (int, string) {
	n, _ := strconv.Atoi(numero)
	if n > 100 {
		return n, "El numero es mayor a 100"
	} else {
		return n, "El numero es menor a 100"
	}
}
