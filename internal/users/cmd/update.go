package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/eduardoraider/go-box/internal/users"
	"github.com/eduardoraider/go-box/pkg/requests"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func updateCmd() *cobra.Command {
	var id int32
	var name string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update user name",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" || id <= 0 {
				log.Println("Please user id and name.")
				os.Exit(1)
			}

			user := users.User{Name: name}

			var body bytes.Buffer
			err := json.NewEncoder(&body).Encode(user)
			if err != nil {
				log.Printf("Error encoding folder %s: %s", name, err)
				os.Exit(1)
			}

			path := fmt.Sprintf("/users/%d", id)
			_, err = requests.AuthenticatedPut(path, &body)
			if err != nil {
				log.Printf("Error updating user %s: %s", name, err)
				os.Exit(1)
			}

			log.Printf("User %s updated!", name)
		},
	}

	cmd.Flags().Int32VarP(&id, "id", "", 0, "User ID")
	cmd.Flags().StringVarP(&name, "name", "n", "", "User name")

	return cmd
}
