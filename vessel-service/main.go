package main

import (
	"fmt"
	"log"
	"os"

	micro "github.com/micro/go-micro/v2"
	pb "github.com/vandong9/learn_go_microservice_1/vessel-service/proto/vessel"
)

/////
const (
	defaultHost = "datastore:27017"
)

// func createDummyData(repo MongoRepository) {
// 	defer repo.Close()
// 	vessels := []*pb.Vessel{
// 		{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
// 	}
// 	for _, v := range vessels {
// 		repo.Create(v)
// 	}
// }

func main() {
	// vessels = []*pb.Vessel{
	// 	&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	// }

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(uri)
	if err != nil {
		log.Panic("")
	}

	vesselCollection := client.Database("shippy").Collection("vessels")
	repo := &MongoRepository{vesselCollection}
	srv := micro.NewService(micro.Name("Vessel.Service"))
	srv.Init()
	// h := &handler{repo}

	pb.RegisterVesselServiceHandler(srv.Server(), &handler{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
