package iteraciones

import (
	"fmt"
)

func Iterar() {

	for i := 0; i < 100; i += 5 {
		if i == 25 {
			continue //salta la iteracion (vuelve a evaluar)
		}
		fmt.Println(i)
		if i == 50 {
			break //corta el ciclo
		}
	}
}
