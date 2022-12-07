package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"José João": 11325.45,
		},
		"P": {
			"Pedro Junior": 1200.00,
		},
	}

	delete(funcsPorLetra, "P") // Excluindo o elemento P do Map

	for _, funcs := range funcsPorLetra {
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
