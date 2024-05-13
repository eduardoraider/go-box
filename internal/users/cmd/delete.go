package cmd

import (
	"fmt"
	"github.com/eduardoraider/go-box/pkg/requests"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func deleteCmd() *cobra.Command {
	var id int32

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete user",
		Run: func(cmd *cobra.Command, args []string) {
			if id < 0 {
				log.Println("ID is required")
				os.Exit(1)
			}

			path := fmt.Sprintf("/users/%d", id)

			err := requests.AuthenticatedDelete(path)
			if err != nil {
				log.Printf("Error deleting user: %v", err)
				os.Exit(1)
			}

			log.Println("User deleted successful")

		},
	}

	cmd.Flags().Int32VarP(&id, "id", "", 0, "User ID")

	return cmd
}
