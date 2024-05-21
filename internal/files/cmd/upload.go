package cmd

import (
	"bytes"
	"github.com/eduardoraider/go-box/pkg/requests"
	"github.com/spf13/cobra"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
)

func uploadCmd() *cobra.Command {
	var folderID int32
	var fileName string

	cmd := &cobra.Command{
		Use:   "upload",
		Short: "Upload a new file",
		Run: func(cmd *cobra.Command, args []string) {
			if fileName == "" {
				log.Println("Patch file name is required")
				os.Exit(1)
			}

			file, err := os.Open(fileName)
			if err != nil {
				log.Printf("Error opening file: %v", err)
				os.Exit(1)
			}
			defer file.Close()

			var body bytes.Buffer

			mw := multipart.NewWriter(&body)
			w, err := mw.CreateFormFile("file", filepath.Base(file.Name()))
			if err != nil {
				log.Printf("Error creating multipart writer: %v", err)
				os.Exit(1)
			}
			_, err = io.Copy(w, file)
			if err != nil {
				log.Printf("Error copying file: %v", err)
				os.Exit(1)
			}

			if folderID > 0 {
				w, err = mw.CreateFormField("folder_id")
				if err != nil {
					log.Printf("Error creating multipart writer: %v", err)
					os.Exit(1)
				}
				w.Write([]byte(strconv.Itoa(int(folderID))))
			}

			mw.Close()

			headers := map[string]string{
				"Content-Type": mw.FormDataContentType(),
			}

			_, err = requests.AuthenticatedPostWithHeaders("/files", &body, headers)
			if err != nil {
				log.Printf("Error uploading file: %v", err)
				os.Exit(1)
			}

			log.Println("File upload successful")
		},
	}

	cmd.Flags().StringVarP(&fileName, "filename", "f", "", "File path")
	cmd.Flags().Int32VarP(&folderID, "folder", "p", 0, "Folder ID")

	return cmd
}
