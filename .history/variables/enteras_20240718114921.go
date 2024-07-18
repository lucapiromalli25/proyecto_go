package variables

import "fmt"

func MostrarEnteros() {
	//variables
	var intcomun int     //por declaracion
	intde32 := int32(10) //declaracion por asiganacion
	intde64 := int64(100)

	fmt.Println("intcomun = ", intcomun)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)
}
