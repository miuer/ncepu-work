package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/miuer/ncepu-work/architecture/com-experiment/soa/weather"
)

const (
	address = "192.168.1.2:10020"
)

// Run -
func Run(location string) (weatherInfo string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	c := pb.NewWeatherClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.QueryWeather(ctx, &pb.QueryRequest{
		Name: location,
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return res.GetMessage()
}
