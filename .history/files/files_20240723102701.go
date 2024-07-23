package files

import (
	//	"bufio"
	"bufio"
	"fmt"

	//"io/ioutil" //Para manejo de archivos (esta deprecated)
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
	if !Append(path, texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(path, texto string) bool {
	arch, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0644) //Abro el archivo para manejarlo, en este caso, de escritura y de append, para que no borre lo anterior
	if err != nil {
		fmt.Println("Error durante el append" + err.Error())
		return false
	}
	_, err = arch.WriteString(texto) //grabar string en el archivo

	if err != nil {
		fmt.Println("Error durante la escritura en el archivo" + err.Error())
		return false
	}

	arch.Close()
	return true
}

func LeoArchivo() {
	/*archivo, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("Error al leer el archivo" + err.Error())
		return
	}

	fmt.Println(string(archivo)) //archivo es un slice de bytes por eso debo convertirlo a string para mostrar por pantalla*/

	archivo, err := os.Open(path)

	if err != nil {
		fmt.Println("Error al leer el archivo" + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo) //los datos los toma del archivo, no del teclado
}
