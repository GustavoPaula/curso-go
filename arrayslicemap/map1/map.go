package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// Map devem ser inicializados
	aprovados := make(map[int]string) // O Map é um array com chave e valor

	aprovados[12345678978] = "Maria" // Chave: 12345678978, Valor: Maria
	aprovados[98765432100] = "Pedro"
	aprovados[94123570623] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[94123570623])
	delete(aprovados, 94123570623) // Deletando o elemento com a chave: 94123570623
	fmt.Println(aprovados)
}
