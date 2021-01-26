// consignment-service/handler.go
package main

import (
	"context"
	"log"

	pb "github.com/vandong9/learn_go_microservice_1/consignment-service/proto/consignment"
	vesselProto "github.com/vandong9/learn_go_microservice_1/vessel-service/proto/vessel"
)

type handler struct {
	repo         repository
	vesselClient vesselProto.VesselService
}

func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	vesselResponse, err := s.vesselClient.FindAvaible(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)

	if err != nil {
		return err
	}
	req.VesselId = vesselResponse.Vessel.Id

	err = s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

func (s *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments, err := s.repo.GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}
