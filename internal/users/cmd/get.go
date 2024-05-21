package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/eduardoraider/go-box/internal/users"
	"github.com/eduardoraider/go-box/pkg/requests"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func getCmd() *cobra.Command {
	var id int32

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get user",
		Run: func(cmd *cobra.Command, args []string) {
			if id < 0 {
				log.Println("ID is required")
				os.Exit(1)
			}

			path := fmt.Sprintf("/users/%d", id)
			data, err := requests.AuthenticatedGet(path)
			if err != nil {
				log.Printf("Error getting user: %v", err)
				os.Exit(1)
			}

			var u users.User
			err = json.Unmarshal(data, &u)
			if err != nil {
				log.Printf("Error unmarshalling user: %v", err)
				os.Exit(1)
			}

			log.Println(u.Name)
			log.Println(u.Login)
			log.Println(u.LastLogin)
		},
	}

	cmd.Flags().Int32VarP(&id, "id", "", 0, "User ID")

	return cmd
}
