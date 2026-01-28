package main

import (
	// "fmt"
	// "runtime"

	"github.com/angorita/aaa/ejercicios"
	"github.com/angorita/aaa/iteraciones"
	// "github.com/angorita/aaa/teclado"
)

func main() {
	// variables.MuestroEnteros()
	// variables.RestoVariables()
	// texto, estado := variables.ConviertoATexto(93)
	// fmt.Println(texto, estado)
	// if os := runtime.GOOS; os == "linux" || os == "Os" {
	// 	fmt.Println("Es Linux")

	// } else {
	// 	fmt.Println("Es windows")
	// }
	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Linux")
	// case "windows":
	// 	fmt.Println("windows")
	// default:
	// 	fmt.Printf("%s\n", os)
	// }
	// num, texto := ejercicios.Convertir("127")
	// fmt.Println(num)
	// fmt.Println(texto)
	// teclado.Ingreso()
	iteraciones.Iterar()
	iteraciones.Break()
	iteraciones.Continue()
	ejercicios.Tabla("9")

}
