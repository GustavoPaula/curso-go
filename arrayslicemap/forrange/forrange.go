package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // O compilador que vai contar!

	for i, elemento := range numeros {
		fmt.Printf("%d) %d\n", i+1, elemento)
	}

	for _, element := range numeros {
		fmt.Println(element)
	}
}
