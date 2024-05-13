package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/eduardoraider/go-box/internal/folders"
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
		Short: "Update folder name",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" || id <= 0 {
				log.Println("Please folder id and name.")
				os.Exit(1)
			}

			folder := folders.Folder{Name: name}

			var body bytes.Buffer
			err := json.NewEncoder(&body).Encode(folder)
			if err != nil {
				log.Printf("Error encoding folder %s: %s", name, err)
				os.Exit(1)
			}

			path := fmt.Sprintf("/folders/%d", id)
			_, err = requests.AuthenticatedPut(path, &body)
			if err != nil {
				log.Printf("Error updating folder %s: %s", name, err)
				os.Exit(1)
			}

			log.Printf("Folder %s updated!", name)
		},
	}

	cmd.Flags().Int32VarP(&id, "id", "n", 0, "Folder ID")
	cmd.Flags().StringVarP(&name, "name", "n", "", "Folder name")

	return cmd
}
