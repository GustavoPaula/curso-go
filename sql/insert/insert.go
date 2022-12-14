package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("insert into usuarios(nome) values(?)") // No insert o ? é o valor que vai ser passado quando chamar a função
	stmt.Exec("Maria")
	stmt.Exec("João")

	res, _ := stmt.Exec("Pedro")

	id, _ := res.LastInsertId() // Pegar o último ID que foi inserido no DB, no caso o ID do Pedro
	fmt.Println(id)

	linhas, _ := res.RowsAffected() // Quantidade de linhas afetadas
	fmt.Println(linhas)
}
