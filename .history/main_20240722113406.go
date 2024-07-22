package main

import (
	"fmt"
	//"runtime"

	//"github.com/lucapiromalli25/proyecto_go/ejercicios"
	//"github.com/lucapiromalli25/proyecto_go/teclado"
	//"github.com/lucapiromalli25/proyecto_go/variables"
	"github.com/lucapiromalli25/proyecto_go/iteraciones"
)

func main() {
	/*fmt.Println("---------------------FUNCION_DOS_DEVOLUCIONES---------------------")
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
	numero, texto := ejercicios.Ejercicio01("100")
	fmt.Println(numero)
	fmt.Println(texto)

	numero2, texto2 := ejercicios.Ejercicio01("200")
	fmt.Println(numero2)
	fmt.Println(texto2)

	numero_err, texto_err := ejercicios.Ejercicio01("HOLA")
	fmt.Println(numero_err)
	fmt.Println(texto_err)

	fmt.Println("---------------------Teclado---------------------")
	teclado.IngresoNumeros()*/

	fmt.Println("---------------------Iteraciones---------------------")
	iteraciones.Iterar()
}
