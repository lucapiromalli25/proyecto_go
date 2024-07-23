package ejercicios

import (
	"bufio" //para lectura por telcado y archivos
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var texto string

func Ejercicio02() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese un número: ")

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
		fmt.Printf("%d . %d = %d \n ", numero, i, numero*i)
	}
}
