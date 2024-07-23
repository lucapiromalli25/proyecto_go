package funciones

import "fmt"

func Calculos() {

	numero3 := 30
	numero4 := 40
	calculo := func(numero1, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}

	fmt.Println(calculo(10, 25))

	calculo = func(numero1, numero2 int) int {
		return numero1 * numero2 * numero3
	}

	fmt.Println(calculo(10, 25))
}
