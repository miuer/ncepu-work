package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// if you want to test this demo in a network model,
	// you should set up address as the following example
	// and set the corresponding port mapping in the gateway
	// 192.168.1.9 -- server IP  :10030 -- any available port

	service := "192.168.1.9:10030"

	// if you want to test this case in a stand-alone(in MAC),
	// you can set up address as the following example
	// 	service := "127.0.0.1:10030"
	// service := "localhost:10030"

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

	fmt.Println(string(buffer))

	sql := `SELECT * FROM students WHERE id = ?`

	var (
		id    string
		name  string
		class string
	)

	db.QueryRow(sql, string(buffer)).Scan(&id, &name, &class)

	str := fmt.Sprintf("学号: %s  姓名: %s  班级: %s ", id, name, class)

	conn.Write([]byte(str))
}
