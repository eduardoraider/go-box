package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/eduardoraider/go-box/internal/folders"
	"github.com/eduardoraider/go-box/pkg/requests"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func listCmd() *cobra.Command {
	var id int32

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List contents of a folder",
		Run: func(cmd *cobra.Command, args []string) {
			path := "/folders"
			if id > 0 {
				path = fmt.Sprintf("/folders/%d", id)
			}

			data, err := requests.AuthenticatedGet(path)
			if err != nil {
				log.Printf("Error getting folders: %v", err)
				os.Exit(1)
			}

			var fc folders.FolderContent
			err = json.Unmarshal(data, &fc)
			if err != nil {
				log.Printf("Error unmarshalling folders: %v", err)
				os.Exit(1)
			}

			log.Println(fc.Folder.Name)
			log.Println("===============")
			for _, c := range fc.Content {
				log.Println(c.ID, " - ", c.Type, " - ", c.Name)
			}
		},
	}

	cmd.Flags().Int32VarP(&id, "id", "", 0, "Folder ID")

	return cmd
}
