package main

import (
	"context"
	"log"
	"os"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/config/cmd"
	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	pb "github.com/vandong9/learn_go_microservice_1/user-service/proto/user"
)

func main() {

	cmd.Init()

	// Create new greeter client
	// client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)
	client := pb.NewUserService("go.micro.srv.user", client.NewClient())

	// flags := micro.Flags(
	// 	cli.StringFlag{
	// 		Name:  "name",
	// 		Usage: "You full name",
	// 	},
	// 	cli.StringFlag{
	// 		Name:  "email",
	// 		Usage: "Your email",
	// 	},
	// 	cli.StringFlag{
	// 		Name:  "password",
	// 		Usage: "Your password",
	// 	},
	// 	cli.StringFlag{
	// 		Name:  "company",
	// 		Usage: "Your company",
	// 	},
	// )
	service := micro.NewService(micro.Name("user-cli"))

	service.Init(

		micro.Action(func(c *cli.Context) error {

			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			company := c.String("company")

			r, err := client.Create(context.TODO(), &pb.User{
				Name:     name,
				Email:    email,
				Password: password,
				Company:  company,
			})
			if err != nil {
				log.Fatalf("Could not create: %v", err)
			}
			log.Printf("Created: %t", r.User.Id)

			getAll, err := client.GetAll(context.Background(), &pb.Request{})
			if err != nil {
				log.Fatalf("Could not list users: %v", err)
			}
			for _, v := range getAll.Users {
				log.Println(v)
			}

			// let's just exit because
			os.Exit(0)
			return err
		}),
	)

	// Run the server
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
