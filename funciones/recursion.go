package funciones

import "fmt"

func Exponencia(valor int) {
	if valor > 1000000 {
		//termina cuando llega a 999999
		return
	}
	fmt.Println("Voy llamando exponencia hasta que llega a 999999", valor)
	Exponencia(valor * 2)
}
func Fibonacci(n int) int {
	// Casos base: si es 0 o 1, retorna el mismo valor
	if n <= 1 {
		return n
	}
	// Llamada recursiva doble
	return Fibonacci(n-1) + Fibonacci(n-2)
}
func FibonacciIterativo(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		// En cada paso: b se vuelve la suma, y a toma el valor anterior de b
		a, b = b, a+b
		fmt.Println(a, b)
	}
	return b
}

/* El código toma un punto de partida y lo duplica en cada iteración hasta que alcanza o supera el millón. Aquí tienes los detalles clave:
Punto de parada: La condición if valor > 1000000 actúa como el caso base, evitando un desbordamiento de pila (stack overflow) [1, 2].
Comportamiento: Si llamas a Exponencia(2), imprimirá potencias de 2 (2, 4, 8, 16...) hasta el 524.288.
Ojo con el cero: Si pasas un 0, entrarás en un bucle infinito (recursión infinita) porque 0 * 2 siempre será 0, nunca superando el millón [2].  */
