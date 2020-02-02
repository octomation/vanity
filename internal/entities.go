package internal

type Module struct {
	Prefix   string   `yaml:"prefix"`
	Import   []Import `yaml:"import"`
	Packages []string `yaml:"packages"`
	Tags     []string `yaml:"tags"`
}

type Import struct {
	URL    string `yaml:"url"`
	VCS    string `yaml:"vcs"`
	Source Source `yaml:"source"`
}

type Source struct {
	URL  string `yaml:"url"`
	Dir  string `yaml:"dir"`
	File string `yaml:"file"`
}

// mapping:
// - prefix -> metaImport.Prefix (cmd/go/internal/get/vcs.go)
// - import.url -> metaImport.RepoRoot (cmd/go/internal/get/vcs.go)
// - import.vcs -> metaImport.VCS (cmd/go/internal/get/vcs.go)
