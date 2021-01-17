package main

import (
	"context"
	"log"
	"sync"

	micro "github.com/micro/go-micro/v2"
	pb "github.com/vandong9/learn_go_microservice_1/consignment-service/proto/consignment"
	vesselProto "github.com/vandong9/learn_go_microservice_1/vessel-service/proto/vessel"
)

const (
	port = ":50051"
)

type repository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type service struct {
	repo         repository
	vesselClient vesselProto.VesselService
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	vesselResponse, err := s.vesselClient.FindAvaible(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)

	if err != nil {
		return err
	}
	req.VesselId = vesselResponse.Vessel.Id

	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {
	repo := &Repository{}

	srv := micro.NewService(micro.Name("service.consignment"))
	srv.Init()

	vesselClient := vesselProto.NewVesselService("vessel.service", srv.Client())

	// pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})
	// pb.RegisterShippingServiceHandler(srv.r(), &service{repo})
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, vesselClient})
	log.Println("Running on port: ", port)
	if err := srv.Run(); err != nil {
		log.Fatalf("fail to serve %v", err)
	}

}
