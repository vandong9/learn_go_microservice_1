package main 

import(
	"sync"
	"context"
	"log"
	"sync"

	pb "github.com/vandong9/learn_go_microservice_1/consignment-service/proto/consignment"
)

const (
	port := ":50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
}

type Repository struct {
	mu 				sync.RWMutex
	consignments 	[]*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()
	return consignment, nil
}

type service struct {
	repo repository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {

}