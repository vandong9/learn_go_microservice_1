// consignment-service/main.go
package main

import (
	"context"
	"log"
	"os"

	micro "github.com/micro/go-micro/v2"
	pb "github.com/vandong9/learn_go_microservice_1/consignment-service/proto/consignment"
	vesselProto "github.com/vandong9/learn_go_microservice_1/vessel-service/proto/vessel"
)

const (
	port        = ":50051"
	defaultHost = "datastore:27017"
)

func main() {
	// repo := &RepositoryMongo{}

	srv := micro.NewService(micro.Name("service.consignment"))
	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}

	defer client.Disconnect(context.TODO())

	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselProto.NewVesselService("vessel.service", srv.Client())
	h := &handler{repository, vesselClient}
	// pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})
	// pb.RegisterShippingServiceHandler(srv.r(), &service{repo})
	pb.RegisterShippingServiceHandler(srv.Server(), h)
	log.Println("Running on port: ", port)
	if err := srv.Run(); err != nil {
		log.Fatalf("fail to serve %v", err)
	}

}
