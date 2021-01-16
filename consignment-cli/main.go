package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	micro "github.com/micro/go-micro"
	pb "github.com/vandong9/learn_go_microservice_1/consignment-service/proto/consignment"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func main() {
	service := micro.NewService(micro.Name("service.consignment.cli"))
	service.Init()

	client := pb.NewShippingServiceClient("service.consignment", service.Client())

	file := defaultFilename

	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not  greet: %v", err)
	}

	log.Printf("Created: %v", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &consignment)
	return consignment, nil
}
