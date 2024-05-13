package main

import (
	authCmd "github.com/eduardoraider/go-box/internal/auth/cmd"
	filesCmd "github.com/eduardoraider/go-box/internal/files/cmd"
	foldersCmd "github.com/eduardoraider/go-box/internal/folders/cmd"
	userCmd "github.com/eduardoraider/go-box/internal/users/cmd"
	"github.com/spf13/cobra"
	"log"
)

var RootCmd = &cobra.Command{}

func main() {
	authCmd.Register(RootCmd)
	userCmd.Register(RootCmd)
	filesCmd.Register(RootCmd)
	foldersCmd.Register(RootCmd)

	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
