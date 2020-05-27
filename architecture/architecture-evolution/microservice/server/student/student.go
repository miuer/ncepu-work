package main

import (
	"context"
	"database/sql"
	"log"
	"net"

	pb "github.com/miuer/ncepu-work/architecture/architecture-evolution/microservice/protobuf/student"
	"google.golang.org/grpc"

	_ "github.com/go-sql-driver/mysql"
)

const (
	network = "tcp"
	address = "192.168.1.9:10020"
)

type studentService struct {
}

func (s *studentService) QueryStudent(ctx context.Context, req *pb.StudentRequest) (resp *pb.StudentResponse, err error) {
	id := req.GetId()

	db, err := sql.Open("mysql", "root:Miufighting.@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sql := `SELECT sname FROM students WHERE id = ?`

	var sname string

	db.QueryRow(sql, id).Scan(&sname)

	resp = &pb.StudentResponse{
		Sname: sname,
	}

	return resp, nil
}

func main() {
	lis, err := net.Listen(network, address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println(address + " net listing...")

	s := grpc.NewServer()

	pb.RegisterStudentServer(s, &studentService{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
