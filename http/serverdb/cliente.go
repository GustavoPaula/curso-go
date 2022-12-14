package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Struct Usuário :)
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// UsuarioHandler analisa o request e delega para função adequada
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/") // Está retirando toda a URL inclusive o /usuarios/, pra ver se tem o ID na URL
	id, _ := strconv.Atoi(sid)                          // Convertendo a string para um valor inteiro

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorID(w, r, id)
	case r.Method == "GET":
		usuarioTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpa... :(")
	}
}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var u Usuario
	db.QueryRow("select id, nome from usuarios where id = ?", id).Scan(&u.ID, &u.Nome) // QueryRow trás apenas uma única linha e Scan mapea os valores para struct

	json, _ := json.Marshal(u) // Convertendo para json

	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, string(json))
}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Fechando a conexão com o banco de dados

	rows, _ := db.Query("select id, nome from usuarios") // Pegando todos os usuários
	defer rows.Close()                                   // Fechando a conexão do select

	var usuarios []Usuario // Criando um Slice do tipo struct Usuario
	for rows.Next() {      // Laço indeterminado, irá percorrer enquanto estiver linhas
		var usuario Usuario                   // Criando uma struct usuario
		rows.Scan(&usuario.ID, &usuario.Nome) // Mapeando os valores do ID e Nome na struct Usuario
		usuarios = append(usuarios, usuario)  // Adicionando cada struct no slice
	}

	json, _ := json.Marshal(usuarios) // Convertendo Struct em json

	w.Header().Set("Content-Type", "application/json") // Enviando a resposta
	fmt.Fprint(w, string(json))                        // Convertendo o json em string
}
