package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//DML - data modification language
	stmt, _ := db.Prepare("Insert into usuarios(nome) values(?)")
	stmt.Exec("Maria")
	stmt.Exec("João")

	res, _ := stmt.Exec("Pedro")
	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()
	fmt.Println(linhas)
}
