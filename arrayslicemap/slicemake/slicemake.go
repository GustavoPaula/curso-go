package main

import "fmt"

func main() {
	s := make([]int, 10) // Criando um slice do tipo int através do make e já definindo 10 elementos
	s[9] = 12            // Atribuindo valor para a última posição do slice
	fmt.Println(s)

	s = make([]int, 10, 20)        // Slice vai ter 10 posições, porém o array interno que foi criado, vai ter 20 posições
	fmt.Println(s, len(s), cap(s)) // len tamanho do slice e o cap é a capacidade máxima do array interno

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) // Append adiciona valores para um slice
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1) // Por padrão quando atinge o tamanho máximo do array interno do slice, ele dropa a capacidade do array.
	fmt.Println(s, len(s), cap(s))
}
