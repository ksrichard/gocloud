package service

import (
	"fmt"
	"github.com/ksrichard/gocloud/model"
	"github.com/ksrichard/gocloud/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var TemplatesYamlFileName = "templates.yaml"

func GetTemplateProjectYaml(templateDir string) (*model.Templates, error) {
	var templateProject model.Templates
	yamlFile, err := ioutil.ReadFile(templateDir + "/" + TemplatesYamlFileName)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &templateProject)
	if err != nil {
		return nil, err
	}
	return &templateProject, nil
}

func GetAllTemplates(projectDir string) ([]model.Template, error) {
	result := []model.Template{}
	dirs, err := ioutil.ReadDir(projectDir)
	if err != nil {
		return nil, err
	}

	for _, d := range dirs {
		if d.IsDir() {
			templateYaml := fmt.Sprintf("%s/%s/template.yaml", projectDir, d.Name())
			if util.FileExists(templateYaml) {
				var tmpl model.Template
				yamlFile, err := ioutil.ReadFile(templateYaml)
				if err != nil {
					return nil, err
				}
				err = yaml.Unmarshal(yamlFile, &tmpl)
				if err != nil {
					return nil, err
				}
				tmpl.Folder = fmt.Sprintf("%s/%s/", projectDir, d.Name())
				result = append(result, tmpl)
			}
		}
	}

	return result, nil
}
