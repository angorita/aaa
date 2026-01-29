package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Tabla() string {
	scanner := bufio.NewScanner(os.Stdin)
	var numero int
	var err error
	var texto string
	for {
		fmt.Println("Ingrese un numero: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}

		}
	}
	fmt.Print("Tabla del : ", numero)
	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("\n%d x %d = %d", numero, i, i*numero)
	}
	return texto
}
