package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()

	//DML - data modification language
	stmt, _ := tx.Prepare("Insert into usuarios(id,nome) values(?,?)")
	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Gustavo")

	_, err = stmt.Exec(1, "Tiago") //chave dublicada
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()

}
