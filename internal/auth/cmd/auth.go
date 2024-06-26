package cmd

import (
	"github.com/eduardoraider/go-box/pkg/requests"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func authenticate() *cobra.Command {
	var (
		user string
		pass string
	)
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate user via API",
		Run: func(cmd *cobra.Command, args []string) {
			if user == "" || pass == "" {
				log.Println("user and pass required")
				os.Exit(1)
			}

			mode := cmd.Parent().Flag("mode").Value.String()

			switch mode {
			case "http":
				authWithHTTP(user, pass)
			case "grpc":
				authWithGRPC(user, pass)
			}

		},
	}

	cmd.Flags().StringVarP(&user, "user", "u", "", "username")
	cmd.Flags().StringVarP(&pass, "pass", "p", "", "password")

	return cmd
}

func authWithHTTP(user, pass string) {
	err := requests.HTTPAuth("/auth", user, pass)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func authWithGRPC(user, pass string) {
	err := requests.GRPCAuth(user, pass)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
