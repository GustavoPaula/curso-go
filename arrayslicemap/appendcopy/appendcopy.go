package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array1 = append(array1, 4, 5, 6)
	slice1 = append(slice1, 4, 5, 6) // Adicionando os valor 4 5 6 para o slice1
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2) // Criando um slice com 2 capacidade
	copy(slice2, slice1)     // Copiando os dados do slice1 para o slice2, porém o copy não aumenta a capacidade do array interno
	fmt.Println(slice2)
}
