package cmd

import (
	"os"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.octolab.org/safe"
	"go.octolab.org/unsafe"
	"gopkg.in/yaml.v2"

	"go.octolab.org/vanity/internal"
	"go.octolab.org/vanity/internal/vanity"
)

func NewDumpCommand() *cobra.Command {
	const (
		file = "modules.yml"
		host = "go.octolab.org"
	)
	command := cobra.Command{
		Use: "dump",
		RunE: func(cmd *cobra.Command, args []string) error {
			file, err := os.Open(file)
			if err != nil {
				return err
			}
			defer safe.Close(file, unsafe.Ignore)

			var modules []internal.Module
			if err := yaml.NewDecoder(file).Decode(&modules); err != nil {
				return err
			}

			wd, err := os.Getwd()
			if err != nil {
				return err
			}
			return vanity.New(host, afero.NewOsFs()).PublishAt(wd, modules)
		},
	}
	return &command
}
