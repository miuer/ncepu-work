package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net"

	pb "github.com/miuer/ncepu-work/architecture/architecture-evolution/microservice/protobuf/grade"
	"google.golang.org/grpc"

	_ "github.com/go-sql-driver/mysql"
)

const (
	network = "tcp"
	address = "192.168.1.9:10021"
)

type gradeService struct {
}

type grade struct {
	id     int
	sid    string
	sname  string
	Course int `json:"course"`
	Score  int `json:"score"`
}

func (s *gradeService) QueryGrade(ctx context.Context, req *pb.GradeRequest) (resp *pb.GradeResponse, err error) {
	id := req.GetId()

	db, err := sql.Open("mysql", "root:Miufighting.@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sql := `SELECT course, score FROM grade WHERE sid = ?;`

	rows, _ := db.Query(sql, id)

	var grades []grade

	for rows.Next() {
		var g grade
		rows.Scan(&g.Course, &g.Score)
		grades = append(grades, g)
	}

	str, _ := json.Marshal(grades)

	resp = &pb.GradeResponse{
		Grade: string(str),
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

	pb.RegisterGradeServer(s, &gradeService{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
