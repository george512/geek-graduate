package main

import (
	"flag"
	"fmt"
	"geek-graduate/internal/client"
	"geek-graduate/pb"
	"geek-graduate/pkg/sample"
	"google.golang.org/grpc"
	"log"
	"strings"
	"time"
)

func testCreateLaptop(laptopClient *client.LaptopClient, laptop *pb.Laptop) {
	laptopClient.CreateLaptop(laptop)
}

func testsearchLaptop(laptopClient *client.LaptopClient) {
	for i := 0; i < 10; i++ {
		laptopClient.CreateLaptop(sample.NewLaptop())
	}

	filter := &pb.Filter{
		MaxPriceUsd: 5000,
		MinCpuCores: 3,
		MinRam:      &pb.Memory{Value: 8, Unit: pb.Memory_GIGABYTE},
		MinCpuGhz:   2.5,
	}

	laptopClient.SearchLaptop(filter)
}

const (
	username        = "admin1"
	password        = "secret"
	refreshDuration = time.Minute * 5
)

func testuploadImage(laptopClient *client.LaptopClient) {
	laptop := sample.NewLaptop()
	laptopClient.CreateLaptop(laptop)
	laptopClient.UploadImage(laptop.GetId(), "tmp/laptop.jpg")
}

func authMethod() map[string]bool {
	const laptopServicePath = "/george.pcbook.CreateLaptop/"
	return map[string]bool{
		laptopServicePath + "CreateLaptop": true,
		laptopServicePath + "UploadImage":  true,
		laptopServicePath + "RateLaptop":   true,
	}
}

func testRateLaptop(laptopClient *client.LaptopClient) {
	n := 3
	laptopIDs := make([]string, n)

	for i := 0; i < n; i++ {
		laptop := sample.NewLaptop()
		laptopIDs[i] = laptop.GetId()
		laptopClient.CreateLaptop(laptop)
	}

	scores := make([]float64, n)

	for {
		fmt.Print("rate laptop (y/n)?  ")
		var answer string
		fmt.Scan(&answer)

		if strings.ToLower(answer) != "y" {
			break
		}

		for i := 0; i < n; i++ {
			scores[i] = sample.RandomLaptopScore()
		}

		err := laptopClient.RateLaptop(laptopIDs, scores)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	serverAddress := flag.String("address", "0.0.0.0:8085", "the server address")
	flag.Parse()

	transportOption := grpc.WithInsecure()
	cc1, err := grpc.Dial(*serverAddress, transportOption)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	authclient := client.NewAuthClient(cc1, username, password)
	interceptor, err := client.NewAuthInterceptor(authclient, authMethod(), refreshDuration)
	if err != nil {
		log.Fatal("cannot create auth interceptor: ", err)
	}

	cc2, err := grpc.Dial(
		*serverAddress,
		transportOption,
		grpc.WithUnaryInterceptor(interceptor.Unary()),
		grpc.WithStreamInterceptor(interceptor.Stream()),
	)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	laptopClient := client.NewLaptopClient(cc2)
	//testuploadImage(laptopClient)
	testRateLaptop(laptopClient)
}
