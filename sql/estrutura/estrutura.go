package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql) // Executando os comandos sql que for passado por parâmetro para essa função
	if err != nil {
		panic(err) // Se der erro, irá mostrar no consolte
	}
	return result // Retornando o resultado
}

func main() {
	db, err := sql.Open("mysql", "root:123456@/") // Abrindo conexão com o DB, mysql(driver de conexão) e root:123456@/(usuário e senha)
	if err != nil {
		panic(err) // Caso houver erro, irá parar o programa
	}
	defer db.Close() // Antes de finalizar a função main, irá fechar a conexão com o banco de dados

	exec(db, "create database if not exists cursogo") // Se o DB não existir, cria o DB com nome cursogo
	exec(db, "use cursogo")                           // Utilizando o DB cursogo
	exec(db, "drop table if exists usuarios")         // Caso existir a tabela usuarios, o Go irá excluir
	exec(db, `create table usuarios(
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY(id)
	)`)
}
