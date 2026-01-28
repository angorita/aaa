package main

import (
	"fmt"

	"github.com/angorita/aaa/variables"
)

func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()
	texto, estado := variables.ConviertoATexto(93)
	fmt.Println(texto, estado)
}
