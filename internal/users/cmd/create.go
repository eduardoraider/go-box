package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/eduardoraider/go-box/internal/users"
	"github.com/eduardoraider/go-box/pkg/requests"
	pb "github.com/eduardoraider/go-box/proto/v1/users"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func createCmd() *cobra.Command {
	var (
		name  string
		login string
		pass  string
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new user",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" || login == "" || pass == "" {
				log.Println("You must specify a name, login and password")
				os.Exit(1)
			}

			mode := cmd.Parent().Flag("mode").Value.String()

			switch mode {
			case "http":
				createWithHTTP(name, login, pass)
			case "grpc":
				createWithGRPC(name, login, pass)
			}

			log.Printf("Created user %s", name)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "User name")
	cmd.Flags().StringVarP(&login, "login", "l", "", "User login")
	cmd.Flags().StringVarP(&pass, "pass", "p", "", "User password")

	return cmd

}

func createWithHTTP(name, login, pass string) {
	user := users.User{
		Name:     name,
		Login:    login,
		Password: pass,
	}

	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(user)
	if err != nil {
		log.Printf("Error encoding user %s: %s", name, err)
		os.Exit(1)
	}

	_, err = requests.Post("/users", &body)
	if err != nil {
		log.Printf("Error creating user: %s", err)
		os.Exit(1)
	}
}

func createWithGRPC(name, login, pass string) {
	user := &pb.UserRequest{
		Name:     name,
		Login:    login,
		Password: pass,
	}

	conn := requests.GetGRPCConn()
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	res, err := client.Create(context.Background(), user)
	if err != nil {
		log.Printf("Error creating user: %s", err)
		os.Exit(1)
	}

	fmt.Printf("User id created: %d", res.User.Id)
}
