package ejercicios

import (
	"bufio" //para lectura por telcado y archivos
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func Ejercicio02() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese un número: ")

	if scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())
		for err != nil {
			continue
		}
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(numero, ".", i, "=", numero*i)
	}
}
