package teclado

import (
	"bufio" //para lectura por telcado y archivos
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {
	fmt.Println("Ingrese número 1: ")

	scanner := bufio.NewScanner(os.Stdin) //Inicializa el objeto scanner, y va a poder leer por teclado

	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
	}
}
