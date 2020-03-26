package vanity

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"

	"go.octolab.org/vanity/internal"
)

var tpl = template.Must(template.New("vanity").Parse(`
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta name="go-import" content="{{ .Package }} {{ .VCS }} {{ .URL }}">
        <meta name="go-source" content="{{ .Package }} {{ .Source.URL }} {{ .Source.Dir }} {{ .Source.File }}">
        <meta http-equiv="content-type" content="text/html; charset=utf-8">
        <meta http-equiv="refresh" content="0; url=https://pkg.go.dev/{{ .Package }}">
        <title>{{ .Package }}</title>
    </head>
    <body>
        Nothing to see here. Please <a href="https://pkg.go.dev/{{ .Package }}">move along</a>.
    </body>
</html>
`))

func New(fs afero.Fs) *publisher {
	return &publisher{fs}
}

type publisher struct {
	fs afero.Fs
}

func (p *publisher) PublishAt(root string, modules []internal.Module) error {
	type Source struct {
		URL, Dir, File string
	}
	type Package struct {
		Package  string
		VCS, URL string
		Source
	}

	for _, module := range modules {
		for _, pkg := range module.Packages {
			dir := filepath.Join(append(
				[]string{root},
				strings.Split(strings.TrimPrefix(pkg, module.Prefix), "/")...,
			)...)
			if err := p.fs.MkdirAll(dir, os.ModePerm); err != nil {
				return err
			}
			file, err := p.fs.Create(filepath.Join(dir, "index.html"))
			if err != nil {
				return err
			}
			if err := tpl.Execute(file, Package{}); err != nil {
				return err
			}
		}
	}
	return nil
}
