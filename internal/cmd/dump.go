package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"go.octolab.org/safe"
	"gopkg.in/yaml.v2"

	"go.octolab.org/vanity/internal"
)

func NewDumpCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "dump",
		RunE: func(cmd *cobra.Command, args []string) error {
			file, err := os.Open("graph.yml")
			if err != nil {
				return err
			}
			defer safe.Close(file, func(err error) { log.Println(err) })

			var modules []internal.Module
			if err := yaml.NewDecoder(file).Decode(&modules); err != nil {
				return err
			}

			cmd.Println(modules)
			return nil
		},
	}
	return cmd
}
