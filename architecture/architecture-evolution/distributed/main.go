package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
)

type grade struct {
	id     int
	sid    string
	sname  string
	Course int `json:"course"`
	Score  int `json:"score"`
}

func main() {

	// server 1
	s1 := "192.168.1.9:10030"

	tcpAddr1, err := net.ResolveTCPAddr("tcp4", s1)
	checkError(err)

	conn1, err := net.DialTCP("tcp", nil, tcpAddr1)
	checkError(err)

	var str []byte
	fmt.Print("please entry your student number:")
	fmt.Scanln(&str)

	conn1.Write(str)

	// server 2
	s2 := "192.168.1.9:10031"

	conn2, err := net.Dial("tcp", s2)
	checkError(err)

	conn2.Write(str)

	// var result []byte
	result := make([]byte, 1024)
	_, err = conn1.Read(result)
	checkError(err)
	sname := string(result)

	n, err := conn2.Read(result)
	checkError(err)

	var grades []grade
	json.Unmarshal(result[:n], &grades)

	for k := range grades {
		fmt.Println(string(str) + " " + sname + " " + strconv.Itoa(grades[k].Course) + " " + strconv.Itoa(grades[k].Score))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
