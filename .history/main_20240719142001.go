package main

import (
	"fmt"
	"runtime"

	"github.com/lucapiromalli25/proyecto_go/ejercicios"
	"github.com/lucapiromalli25/proyecto_go/variables"
)

func main() {
	Estado, Texto := variables.ConvietoATexto(1231)
	fmt.Println(Estado)
	fmt.Println(Texto)

	fmt.Println("---------------------IF---------------------")
	//Ve en que S.O esta parado
	if os := runtime.GOOS; os == "linux" || os == "OS X." { //puedo asignar variables y preguntarlas en el if
		fmt.Println("Estoy en Linux o Mac, esto es: ", os)
	} else {
		fmt.Println("Estoy en Windows")
	}

	fmt.Println("---------------------SWITCH---------------------")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os) //Printf nos deja formatear el texto de la manera que queramos
	}

	fmt.Println("---------------------Ejercicio1---------------------")
	Numero, Texto := ejercicios.Ejercicio01("100")
	fmt.Println(Numero)
	fmt.Println(Texto)

	Numero2, Texto2 := ejercicios.Ejercicio01("200")
	fmt.Println(Numero2)
	fmt.Println(Texto2)
}
