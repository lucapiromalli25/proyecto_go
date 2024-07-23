package files

import (
	//	"bufio"
	"fmt"
	//	"ioutil" //Para manejo de archivos
	"os"

	"github.com/lucapiromalli25/proyecto_go/ejercicios"
)

var path string = "./files/txt/tabla.txt"

func GrabarTabla() {
	texto := ejercicios.Ejercicio02()

	//creo el archivo
	archivo, err := os.Create(path)
	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto) //Grabo el texto en el archivo
	archivo.Close()
}

func SumaTabla() {
	texto := ejercicios.Ejercicio02()
	if Append(path, texto) == false {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(path, texto string) bool {
	arch, err := os.OpenFile(path, os.O_WDONLY || os.O_APPEND) //Abro el archivo para manejarlo, en este caso, de escritura y de append, para que no borre lo anterior
	return false
}
