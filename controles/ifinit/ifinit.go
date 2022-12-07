package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano()) // Data atual e o nano segundo
	r := rand.New(s)
	return r.Intn(10) // Vai gerar um número aleatório até 10
}

func main() {
	if i := numeroAleatorio(); i > 5 { // Também é suportado no switch e essa variável é visível apenas dentro do IF/ELSE
		fmt.Println("Ganhou!!")
	} else {
		fmt.Println("Perdeu!!")
	}
}
