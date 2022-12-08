package main

import "runtime/debug"

func f3() {
	debug.PrintStack() // Exibe a pilha de execução de um programa no momento que essa função for chamada
}

func f2() {
	f3()
}

func f1() {
	f2()
}

func main() {
	f1()
}
