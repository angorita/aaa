package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1, numero2 int
var leyenda string
var err error

func Ingreso() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese numero 1: ")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Hubo un error" + err.Error())
		}
	}

	fmt.Println("Ingrese numero 2: ")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Hubo un error" + err.Error())
		}
	}

	fmt.Println("Leyenda... ")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}
	fmt.Println(leyenda, numero1*numero2)

}
