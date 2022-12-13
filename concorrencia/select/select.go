package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
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

func oMaisRapido(url1, url2, url3 string) string {
	c1 := titulo(url1)
	c2 := titulo(url2)
	c3 := titulo(url3)

	// estrutura de controle específica para concorrência
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
	default:
		return "Sem resposta ainda!"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.youtube.com",
		"https://www.uol.com.br",
		"https://www.google.com.br",
	)
	fmt.Println(campeao)
}
