package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public")) // Criando um FileServer e o diretório apartir da pasta public
	http.Handle("/", fs)                      // Handle diz que quando chegar um request no "/", passa para o fs(File Server) criado

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil)) // Servidor escutando na porta 6000
}
