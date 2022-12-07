package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { // Quando você não coloca nenhuma condição, ele vai procurar um case que seja true
	case t.Hour() < 2:
		fmt.Println("Bom dia")
	case t.Hour() < 18:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite")
	}
}
