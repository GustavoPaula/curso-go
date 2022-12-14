package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {
	c := make(chan string) // Criando um canal de string
	go func() {            // função anônima
		for i := 0; i < 3; i++ { // For repetindo 3 vezes
			time.Sleep(time.Second)                       // Pause de 1 segundo
			c <- fmt.Sprintf("%s falando: %d", pessoa, i) // Enviando uma string para o canal
		}
	}() // chamando a função anônima
	return c // returnando o canal
}

func juntar(entrada1, entrada2 <-chan string) <-chan string { // recebendo dois canais de string e retornando um canal de string
	c := make(chan string) // Criando um canal de string
	go func() {            // Função anônima
		for { // For infinito
			select { // Select pegando os casos do canal
			case s := <-entrada1: // Caso chega informação no canal1
				c <- s // Pegando o resultado da entrada e atribuindo ao canal
			case s := <-entrada2:
				c <- s
			}
		}
	}() // Chamando a função anônima
	return c // Returnando o canal de string
}

func main() {
	c := juntar(falar("João"), falar("Maria"))
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
}
