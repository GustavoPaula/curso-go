package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

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

func encaminhar(origem <-chan string, destino chan string) <-chan string {
	for {
		destino <- <-origem // Todo canal que receber no origem, irá encaminhar os dados para o canal de destino
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)     // Criando Canal de string
	go encaminhar(entrada1, c) // Passando o canal entrada para o canal c
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		titulo("https://www.google.com", "https://www.cod3r.com.br"),
		titulo("https://www.uol.com", "https://www.youtube.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
