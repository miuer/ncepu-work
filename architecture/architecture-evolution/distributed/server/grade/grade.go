package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type grade struct {
	id     int
	sid    string
	sname  string
	Course int `json:"course"`
	Score  int `json:"score"`
}

func main() {

	service := "192.168.1.9:10031"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	db, err := sql.Open("mysql", "root:Miufighting.@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handle(conn, db)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func handle(conn net.Conn, db *sql.DB) {

	defer conn.Close()

	buffer := make([]byte, 12)
	//	var str []byte
	conn.Read(buffer)

	sql := `SELECT course, score FROM grade WHERE sid = ?;`

	rows, _ := db.Query(sql, string(buffer))

	var grades []grade

	for rows.Next() {
		var g grade
		rows.Scan(&g.Course, &g.Score)
		grades = append(grades, g)
	}

	str, _ := json.Marshal(grades)

	conn.Write(str)

}
