package main

import (
	"context"

	"gopkg.in/mgo.v2/bson"

	pb "github.com/vandong9/learn_go_microservice_1/vessel-service/proto/vessel"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(vessel *pb.Vessel) error
	// Close()
}

// type VesselRepository struct {
// 	vessels []*pb.Vessel
// }

type MongoRepository struct {
	collection *mongo.Collection
}

/// FindAvailable
func (repo *MongoRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {

	var vessel *pb.Vessel
	err := repo.collection.FindOne(context.Background(), bson.M{
		"capacity":  spec.Capacity,
		"maxweight": spec.MaxWeight,
	}).Decode(&vessel)
	// err := repo.collection.FindOne(context.Background(), bson.M{
	// 	"capacity":  bson.M{"$gte": spec.Capacity},
	// 	"maxweight": bson.M{"$gte": spec.MaxWeight},
	// }).Decode(&vessel)
	//.One(&vessel)
	if err != nil {
		return nil, err
	}
	return vessel, nil

}

func (repository *MongoRepository) Create(vessel *pb.Vessel) error {
	var _, err = repository.collection.InsertOne(context.Background(), vessel)
	return err
}

// func (repo *MongoRepository) Close() {
// 	repo.session.Close()
// }

// func (repo *MongoRepository) collection() *mgo.Collection {
// 	return repo.session.DB(dbName).C(vesselCollection)
// }
