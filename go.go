package main

func main() {}

type Module struct {
	Prefix string   `yaml:"prefix"`
	Import Import   `yaml:"import"`
	Tags   []string `yaml:"tags"`
	Skip   bool     `yaml:"skip"`
}

type Import struct {
	URL string `yaml:"url"`
	VCS string `yaml:"vcs"`
}
