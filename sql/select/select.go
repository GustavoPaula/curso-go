package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios where id > ?", 5) // Selecionando os usuarios com ID maior que 5
	defer rows.Close()

	for rows.Next() { // For indeterminado, enquanto estiver registro, o laço vai ser executado
		var u usuario             // Criando uma struct
		rows.Scan(&u.id, &u.nome) // Está mapeando o valor da consulta (select) para dentro da struct u
		fmt.Println(u)
	}
}
