package cmd

import (
	"bytes"
	"encoding/json"
	"github.com/eduardoraider/go-box/internal/folders"
	"github.com/eduardoraider/go-box/pkg/requests"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func createCmd() *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new folder",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" {
				log.Println("Please provide a folder name")
				os.Exit(1)
			}

			folder := folders.Folder{Name: name}

			var body bytes.Buffer
			err := json.NewEncoder(&body).Encode(folder)
			if err != nil {
				log.Printf("Error encoding folder %s: %s", name, err)
				os.Exit(1)
			}

			_, err = requests.AuthenticatedPost("/folders", &body)
			if err != nil {
				log.Printf("Error creating folder %s: %s", name, err)
				os.Exit(1)
			}

			log.Printf("Created folder %s", name)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Folder name")

	return cmd
}
