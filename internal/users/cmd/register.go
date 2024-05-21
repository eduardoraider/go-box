package cmd

import "github.com/spf13/cobra"

func Register(c *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "users",
		Short: "Manage users",
	}

	cmd.AddCommand(createCmd())
	cmd.AddCommand(getCmd())
	cmd.AddCommand(listCmd())
	cmd.AddCommand(updateCmd())
	cmd.AddCommand(deleteCmd())

	c.AddCommand(cmd)
}
