package arreglos_slices

import (
	"fmt"
)

var tabla []int = []int{2, 3, 4}
var arreglo [10]int = [10]int{4, 5, 94, 50, 52, 65}

func Slice() {
	// elementos := make([]int, 500, 20000)
	// fmt.Printf("Largo %d,Capacidad %d", len(elementos), cap(elementos))
	num := make([]int, 0, 0)
	for i := 0; i < 9; i++ {
		num = append(num, i)
	}
	fmt.Printf("\nLargo %d,Capacidad %d", len(num), cap(num))

}
func River() {
	// Definimos el map: la clave es el número (int) y el valor el nombre (string)
	plantelRiver := map[int]string{
		1:  "Franco Armani",
		17: "Paulo Díaz",
		10: "Manuel Lanzini",
		9:  "Miguel Borja",
	}

	// Agregar un jugador nuevo
	plantelRiver[11] = "Facundo Colidio"
	plantelRiver[12] = "Oscar Angarita"

	// Acceder a un valor
	fmt.Println("\nEl 12 de River es:", plantelRiver[12])

	// Recorrer el "mapa" de jugadores
	for numero, nombre := range plantelRiver {
		fmt.Printf("Jugador  %d: %s\n", numero, nombre)
	}
}
