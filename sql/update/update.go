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

	//DML - data modification language
	stmt, _ := db.Prepare("update usuarios set nome = ? where id =? ")
	stmt.Exec("Beatriz", 2000)
	stmt.Exec("Gu", 2001)

	stmt2, _ := db.Prepare("delete from usuarios where id =? ")
	stmt2.Exec(3)
}
