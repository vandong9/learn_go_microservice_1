package main

import (
	"context"
	"log"
	"sync"

	micro "github.com/micro/go-micro"
	pb "github.com/vandong9/learn_go_microservice_1/consignment-service/proto/consignment"
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
	repo repository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments := s.repo.GetAll()

	return &pb.Response{Consignments: consignments}, nil
}

func main() {
	repo := &Repository{}

	// lis, err := net.Listen("tcp", port)
	// if err != nil {
	// 	log.Fatalf("fail to listen: %v", err)
	// }

	// s := grpc.NewServer()
	// pb.RegisterShippingServiceServer(s, &service{repo})
	// reflection.Register(s)

	srv := micro.NewService(micro.Name("service.consignment"))
	srv.Init()
	// pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})

	log.Println("Running on port: ", port)
	if err := srv.Run(); err != nil {
		log.Fatalf("fail to serve %v", err)
	}

}
