package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Miufighting.@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	str := `INSERT INTO students(id,sname,class) VALUES("201709001007","苏","网络1701");
	`
	db.Exec(str)
}
