package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

// <-chan - canal somente-leitura
func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url) // Requisição get para cada URL
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>") // Pegando apenas o título da URL
			c <- r.FindStringSubmatch(string(html))[1]       // Convertendo o html em string, pegando o segundo valor no array e gravando no canal
		}(url) // Passando Url como parâmetro para func anônima e já invocando
	}
	return c
}

func main() {
	t1 := titulo("https://www.google.com", "https://www.uol.com.br")
	t2 := titulo("https://www.youtube.com", "https://www.cod3r.com.br")
	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}
