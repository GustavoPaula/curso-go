package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	// fmt.Println("O valor de x é " + x)
	xs := fmt.Sprint(x) // Transforma o número em string
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)

	fmt.Printf("O valor de x é %.2f", x) // Print formatado %s(string) %f(float), o resultado será duas casas decimais depois da vírgula.

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d) // \n quebra a linha, %d(inteiro), %f(float), %.1f(float com 1 casa decimal), %t(boolean), %s(string)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)         // %v é um tipo mais genérico.
}
