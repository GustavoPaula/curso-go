package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Irá fechar a conexão com o banco antes da função main ser finalizada

	tx, _ := db.Begin()                                                  // Inicializar uma transação e retornar o objeto tx(responsável por executar os comandos sql)
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?, ?)") // Retorna statement e um erro

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")
	_, err = stmt.Exec(1, "Tiago") // chave duplicada

	if err != nil {
		tx.Rollback()  // Voltar a transação
		log.Fatal(err) // Mostra o erro
	}

	tx.Commit() // Caso dê certo, irá salvar a transação
}
