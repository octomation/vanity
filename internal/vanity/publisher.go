package vanity

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"

	"go.octolab.org/vanity/internal"
)

func New(host string, fs afero.Fs) *publisher {
	return &publisher{host, fs}
}

type publisher struct {
	host string
	fs   afero.Fs
}

func (p *publisher) PublishAt(root string, modules []internal.Module) error {
	for _, module := range modules {
		if len(module.Import) == 0 {
			continue
		}

		for _, pkg := range module.Packages {
			dir := filepath.Join(append(
				[]string{root},
				strings.Split(strings.TrimPrefix(pkg, module.Name), "/")...,
			)...)
			if err := p.fs.MkdirAll(dir, os.ModePerm); err != nil {
				return err
			}
			file, err := p.fs.Create(filepath.Join(dir, "index.html"))
			if err != nil {
				return err
			}

			// TODO:rfc:#3 add support multiple source
			for _, imports := range module.Import[:1] {
				if err := tpl.Execute(file, Meta{
					Package: pkg,
					Import: MetaImport{
						Prefix:   module.Name,
						VCS:      imports.VCS,
						RepoRoot: imports.URL,
					},
					Source: MetaSource{
						URL:  imports.Source.URL,
						Dir:  imports.Source.Dir,
						File: imports.Source.File,
					},
				}); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
