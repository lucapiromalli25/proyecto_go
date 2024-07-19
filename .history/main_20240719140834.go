package main

import (
	"fmt"
	"runtime"

	"github.com/lucapiromalli25/proyecto_go/variables"
)

func main() {
	Estado, Texto := variables.ConvietoATexto(1231)
	fmt.Println(Estado)
	fmt.Println(Texto)

	//Ve en que S.O esta parado
	if os := runtime.GOOS; os == "linux" || os == "OS X." { //puedo asignar variables y preguntarlas en el if
		fmt.Println("Estoy en Linux o Mac, esto es: ", os)
	} else {
		fmt.Println("Estoy en Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os) //Printf nos deja formatear el texto de la manera que queramos
	}
}
