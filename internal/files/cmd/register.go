package cmd

import "github.com/spf13/cobra"

func Register(c *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "file",
		Short: "Manage files",
	}

	cmd.AddCommand(uploadCmd())
	cmd.AddCommand(updateCmd())
	cmd.AddCommand(deleteCmd())

	c.AddCommand(cmd)
}
