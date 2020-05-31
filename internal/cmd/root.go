package cmd

import "github.com/spf13/cobra"

// New returns the new root command.
func New() *cobra.Command {
	command := cobra.Command{
		Use:   "vanity",
		Short: "Vanity URL manager",
		Long:  "Vanity URL manager.",

		SilenceErrors: false,
		SilenceUsage:  true,
	}
	command.AddCommand(NewDumpCommand())
	return &command
}
