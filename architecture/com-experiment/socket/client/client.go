package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// network model
	// LAN: host IP + port or gateway IP + port
	// WAN: gateway IP + port
	service := "192.168.1.9:10030"
	// service := "222.221.83.170:10030"

	// stand-alone
	// service := "localhost:10030"
	// service := "127.0.0.1:10030"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	var str []byte
	fmt.Print("please entry your student number:")
	fmt.Scanln(&str)

	conn.Write(str)

	// var result []byte
	result := make([]byte, 1024)
	_, err = conn.Read(result)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
