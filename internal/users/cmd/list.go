package cmd

import (
	"encoding/json"
	"github.com/eduardoraider/go-box/internal/users"
	"github.com/eduardoraider/go-box/pkg/requests"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func listCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List users",
		Run: func(cmd *cobra.Command, args []string) {
			data, err := requests.AuthenticatedGet("/users")
			if err != nil {
				log.Printf("Error getting users: %v", err)
				os.Exit(1)
			}

			var us []users.User
			err = json.Unmarshal(data, &us)
			if err != nil {
				log.Printf("Error unmarshalling users: %v", err)
				os.Exit(1)
			}

			for _, u := range us {
				log.Println(u.Name, u.Login, u.LastLogin)
			}
		},
	}

	return cmd
}
