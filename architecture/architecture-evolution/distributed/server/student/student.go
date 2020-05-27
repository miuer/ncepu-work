package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	service := "192.168.1.9:10030"

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

	// fmt.Println(string(buffer))

	sql := `SELECT sname FROM students WHERE id = ?`

	var sname string

	db.QueryRow(sql, string(buffer)).Scan(&sname)

	conn.Write([]byte(sname))
}
