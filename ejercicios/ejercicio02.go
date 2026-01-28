package ejercicios

import (
	"fmt"
	"strconv"
)

func Tabla(numero string) {
	num, err := strconv.Atoi(numero)
	if err != nil {
		fmt.Println("Error en el ingreso de numero" + err.Error())
	}
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d %d", num, i*num)
	}
}
