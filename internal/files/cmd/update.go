package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/eduardoraider/go-box/internal/files"
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
		Short: "Update file name",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" || id <= 0 {
				log.Println("Please file id and name.")
				os.Exit(1)
			}

			file := files.File{Name: name}

			var body bytes.Buffer
			err := json.NewEncoder(&body).Encode(file)
			if err != nil {
				log.Printf("Error encoding file %s: %s", name, err)
				os.Exit(1)
			}

			path := fmt.Sprintf("/files/%d", id)
			_, err = requests.AuthenticatedPut(path, &body)
			if err != nil {
				log.Printf("Error updating file %s: %s", name, err)
				os.Exit(1)
			}

			log.Printf("File %s updated!", name)
		},
	}

	cmd.Flags().Int32VarP(&id, "id", "n", 0, "File ID")
	cmd.Flags().StringVarP(&name, "name", "n", "", "File name")

	return cmd
}
