package files

import (
	//	"bufio"
	"fmt"
	//	"ioutil" //Para manejo de archivos
	"os"

	"github.com/lucapiromalli25/proyecto_go/ejercicios"
)
path:="./files/txt/tabla.txt"
func GrabarTabla() {
	texto := ejercicios.Ejercicio02(path)

	//creo el archivo
	archivo, err := os.Create("./files/txt/tabla.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto) //Grabo el texto en el archivo
	archivo.Close()
}


func SumaTabla() {
	texto := ejercicios.Ejercicio02()
	filename := path
}
