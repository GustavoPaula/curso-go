package main

import (
	"fmt"
	"time"
)

// Channel (canal) - É a forma de comunicação entre gotourines, ele é síncrono, é o ponto de sincronismo
// channel é um tipo (int, float, string)

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)
	fmt.Println("A")

	a, b := <-c, <-c // recebendo os dois primeiros dados do canal
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<-c)
}
