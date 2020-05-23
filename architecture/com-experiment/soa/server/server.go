package main

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

	pb "github.com/miuer/ncepu-work/architecture/com-experiment/soa/weather"
	"google.golang.org/grpc"
)

const (
	network = "tcp"
	address = "192.168.1.9:10020"
)

type weatherService struct {
}

func (s *weatherService) QueryWeather(ctx context.Context, req *pb.QueryRequest) (*pb.QueryReply, error) {

	location := req.GetName()

	httpClient := &http.Client{}

	log.Println(location)

	httpReq, _ := http.NewRequest(
		"POST",
		"http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName",
		strings.NewReader("theCityName="+location),
	)

	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	httpReq.Header.Set("Host", "www.webxml.com.cn")

	httpResp, _ := httpClient.Do(httpReq)

	body, _ := ioutil.ReadAll(httpResp.Body)

	defer httpResp.Body.Close()

	res := pb.QueryReply{
		Message: string(body),
	}

	return &res, nil
}

func main() {
	lis, err := net.Listen(network, address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println(address + " net listing...")

	s := grpc.NewServer()
	pb.RegisterWeatherServer(s, &weatherService{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
