// consignment-service/main.go
package main

import (
	"context"
	"log"
	"os"

	"github.com/micro/go-micro/metadata"
	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/server"
	pb "github.com/vandong9/learn_go_microservice_1/consignment-service/proto/consignment"
	userService "github.com/vandong9/learn_go_microservice_1/user-service/proto/user"
	vesselProto "github.com/vandong9/learn_go_microservice_1/vessel-service/proto/vessel"
)

const (
	port        = ":50051"
	defaultHost = "localhost:27017"
)

func main() {
	// repo := &RepositoryMongo{}

	srv := micro.NewService(micro.Name("service.consignment"), micro.Version("latest"), micro.WrapHandler(AuthWrapper))
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

// AuthWrapper is a high-order function which takes a HandlerFunc
// and returns a function, which takes a context, request and response interface.
// The token is extracted from the context set in our consignment-cli, that
// token is then sent over to the user service to be validated.
// If valid, the call is passed along to the handler. If not,
// an error is returned.
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		if os.Getenv("DISABLE_AUTH") == "true" {
			return fn(ctx, req, resp)
		}
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			log.Println("Missing with token: ")
			// return errors.New("no auth meta-data found in request")
		}
		// Note this is now uppercase (not entirely sure why this is...)
		token := meta["Token"]
		log.Println("Authenticating with token: ", token)
		// Auth here
		authClient := userService.NewUserService("UserService", client.DefaultClient)
		_, err := authClient.ValidateToken(context.Background(), &userService.Token{
			Token: token,
		})
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}
