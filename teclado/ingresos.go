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
	scanner := bufio.NewScanner(os.Stdin) //Inicializa el objeto scanner, y va a poder leer por teclado

	fmt.Println("Ingrese número 1: ")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}
	}

	fmt.Println("Ingrese número 2: ")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda: ")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numero1*numero2)
}
