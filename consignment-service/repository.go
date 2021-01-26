// consignment-service/repository.go
package main

import (
	"context"

	pb "github.com/vandong9/learn_go_microservice_1/consignment-service/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository interface {
	Create(consignment *pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
}

// type Repository struct {
// 	mu           sync.RWMutex
// 	consignments []*pb.Consignment
// }

// func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
// 	repo.mu.Lock()
// 	updated := append(repo.consignments, consignment)
// 	repo.consignments = updated
// 	repo.mu.Unlock()
// 	return nil
// }

// func (repo *Repository) GetAll() []*pb.Consignment {
// 	return repo.consignments
// }

// MongoRepository implementation
type MongoRepository struct {
	collection *mongo.Collection
}

func (repository *MongoRepository) Create(consignment *pb.Consignment) error {
	_, err := repository.collection.InsertOne(context.Background(), consignment)
	return err
}

func (repository *MongoRepository) GetAll() ([]*pb.Consignment, error) {
	cur, err := repository.collection.Find(context.Background(), nil, nil)
	var consignments []*pb.Consignment
	for cur.Next(context.Background()) {
		var consignment *pb.Consignment
		if err := cur.Decode(&consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}
	return consignments, err

	// var consignements []*pb.Consignement
	// for cur.Next(context.Background()) {
	// 	var consignement *pb.Consignment
	// 	if err := cur.Decode(&consignment); err != nil {
	// 		return nil, err
	// 	}
	// 	consignements = append(consignments, consignement)
	// }

	// return consignments, err
}
