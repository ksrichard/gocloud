package model

type Templates struct {
	Projects []TemplateProject `yaml:"projects"`
}

type TemplateProject struct {
	Name   string `yaml:"name"`
	Title  string `yaml:"title"`
	Folder string `yaml:"folder"`
}

type Template struct {
	Folder      string
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	// TODO: add other fields
}
