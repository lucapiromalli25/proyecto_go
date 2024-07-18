package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	//MostrarEnteros() // es visible en cualquier parte del paquete variables
	Nombre = "Luca"
	Estado = true
	Sueldo = 1554.84
	Fecha = time.Now() // time.Now() es una funcion que devuelve la fecha actual

	//El println puede printear cualquier valor porque acepta por parametro any tipo
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)

}

func ConvietoATexto(numero int) (bool, string) {
	var txt string
	txt = strconv.Itoa(numero)
	return true, txt
}
