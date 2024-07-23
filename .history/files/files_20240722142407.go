package files

import (
	"fmt" //Para manejo de archivos
	"os"

	"github.com/lucapiromalli25/proyecto_go/ejercicios"
)

func GrabarTabla() {
	texto := ejercicios.Ejercicio02()

	//creo el archivo
	archivo, err := os.Create("./files/txt/tabla.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo", err)
	}
}

func SumaTabla() {
	texto := ejercicios.Ejercicio02()
}
