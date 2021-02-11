package main

import (
	"context"
	"log"

	"github.com/micro/cli/v2"
	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/config/cmd"
	pb "github.com/vandong9/learn_go_microservice_1/user-service/proto/user"
)

func main() {
	cmd.Init()

	// Create new greeter client
	mclient := pb.NewUserService("go.micro.srv.user", client.DefaultClient)

	// flags := []cli.Flag{
	// 	cli.StringFlag{
	// 		Name:  "lang, l",
	// 		Value: "english",
	// 		Usage: "language for the greeting",
	// 	},
	// }
	// []cli.Flag{cli.StringFlag{Name: "name", Usage: "your full name"},
	// 	cli.StringFlag{Name: "email", Usage: "your email"},
	// 	cli.StringFlag{Name: "password", Usage: "your password"},
	// 	cli.StringFlag{Name: "company", Usage: "your company"}}
	service := micro.NewService(micro.Name("user-service"))

	// Start as service
	service.Init(micro.Action(func(c *cli.Context) error {
		name := c.String("name")
		email := c.String("email")
		password := c.String("password")
		company := c.String("company")

		r, err := mclient.Create(context.TODO(), &pb.User{Name: name, Email: email, Password: password, Company: company})
		if err != nil {
			log.Fatalf("Could not create: %v", err)
			return err
		}
		log.Printf("Created %s", r.User.Id)

		getAll, err := mclient.GetAll(context.Background(), &pb.Request{})
		if err != nil {
			log.Fatalf("Could not list users: %v", err)
			return err
		}

		for _, v := range getAll.Users {
			log.Println(v)
		}
		// os.Exit(0)
		return nil
	}))
}
