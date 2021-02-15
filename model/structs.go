package model

import "github.com/manifoldco/promptui"

type Templates struct {
	Projects []TemplateProject `yaml:"projects"`
}

type TemplateProject struct {
	Name   string `yaml:"name"`
	Title  string `yaml:"title"`
	Folder string `yaml:"folder"`
}

type Template struct {
	Folder         string
	Name           string             `yaml:"name"`
	Description    string             `yaml:"description"`
	Requirements   []string           `yaml:"requirements"`
	InstallScripts []TemplateScript   `yaml:"install_script"`
	RunScripts     []TemplateScript   `yaml:"run_script"`
	Properties     []TemplateProperty `yaml:"properties"`
	PropertyValues map[string]interface{}
}

type TemplateScript struct {
	OS     string `yaml:"os"`
	Script string `yaml:"script"`
}

type TemplateProperty struct {
	Name        string   `yaml:"name"`
	Type        string   `yaml:"type"`
	Description string   `yaml:"description"`
	Required    bool     `yaml:"required"`
	Requires    []string `yaml:"requires"`
	Prompt      *promptui.Prompt
}
