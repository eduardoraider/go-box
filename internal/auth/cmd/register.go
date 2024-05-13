package cmd

import "github.com/spf13/cobra"

func Register(c *cobra.Command) {
	c.AddCommand(authenticate())
}
