package cmd

import "github.com/spf13/cobra"

// New returns the new root command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "vanity",
		Short:         "Vanity URL manager",
		Long:          "Vanity URL manager.",
		SilenceErrors: false,
		SilenceUsage:  true,
	}
	cmd.AddCommand(NewDumpCommand())
	return cmd
}
