package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y)) // Ou fmt.Println(int(x) / y)

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado...
	fmt.Println("Teste " + string(97))

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// string para boolean
	b, _ := strconv.ParseBool("true") // True ou 1 são os valores verdadeiros, os números diferente de 1 é false
	if b {
		fmt.Println("Verdadeiro")
	}
}
