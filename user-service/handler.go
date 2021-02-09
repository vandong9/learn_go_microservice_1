package main

import (
	"context"
	pb "github.com/vandong9/learn_go_microservice_1/user-service/proto/user"
)


type service struct {
	repo         Repository
	tokenService Authable
}


func (srv *service) Get(ctx context.Context, req *pb.User)