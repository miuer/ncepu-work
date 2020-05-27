package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"

	gpb "github.com/miuer/ncepu-work/architecture/architecture-evolution/microservice/protobuf/grade"
	spb "github.com/miuer/ncepu-work/architecture/architecture-evolution/microservice/protobuf/student"
)

const (
	saddress = "192.168.1.9:10020"
	gaddress = "192.168.1.9:10021"
)

type grade struct {
	id     int
	sid    string
	sname  string
	Course int `json:"course"`
	Score  int `json:"score"`
}

func main() {

	var id string
	fmt.Print("please entry your student number:")
	fmt.Scanln(&id)

	sconn, err := grpc.Dial(saddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer sconn.Close()

	sc := spb.NewStudentClient(sconn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	sresp, err := sc.QueryStudent(ctx, &spb.StudentRequest{
		Id: id,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	gconn, err := grpc.Dial(gaddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer gconn.Close()

	gc := gpb.NewGradeClient(gconn)

	gresp, err := gc.QueryGrade(ctx, &gpb.GradeRequest{
		Id: id,
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// log.Println(sresp.GetSname())

	// log.Println(gresp.GetGrade())
	var grades []grade
	json.Unmarshal([]byte(gresp.GetGrade()), &grades)

	for k := range grades {
		fmt.Println(id + " " + sresp.GetSname() + " " + strconv.Itoa(grades[k].Course) + " " + strconv.Itoa(grades[k].Score))
	}
}
