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

	os := runtime.GOOS //Ve en que S.O esta parado
	if os == "Linux." {

	} else {

	}
}
