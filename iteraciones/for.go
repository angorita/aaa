package iteraciones

import "fmt"

func Iterar() {
	for i := 0; i < 10; i += 5 {
		fmt.Println(i)
	}

}
func Break() {
	for i := 10; i > 1; i-- {
		if i == 6 {
			break
		}
		fmt.Println(i)
	}

}
func Continue() {
	for i := 10; i > 1; i-- {
		if i == 6 {
			continue

		}
		fmt.Println(i)
	}

}
