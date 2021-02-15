package service

import (
	"fmt"
	"github.com/ksrichard/gocloud/model"
	"github.com/ksrichard/gocloud/util"
	"github.com/manifoldco/promptui"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
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
			templateYaml := fmt.Sprintf("%s/%s/project.yaml", projectDir, d.Name())
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

func GetProjectYaml(projectDir string) (*model.Template, error) {
	var result model.Template
	templateYaml := fmt.Sprintf("%s/project.yaml", projectDir)
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
		tmpl.Folder = fmt.Sprintf("%s/", projectDir)
		result = tmpl
	}
	return &result, nil
}

func ValidateRequired(input string, required bool) error {
	if required && strings.TrimSpace(input) == "" {
		return util.BoldError("This property is required!")
	}
	return nil
}

func PopulatePrompts(template *model.Template) error {
	for i, property := range template.Properties {
		currentRequired := property.Required
		switch strings.ToLower(property.Type) {
		case "boolean":
			template.Properties[i].Prompt = &promptui.Prompt{
				Label:     util.Bold().Sprint(property.Description),
				IsConfirm: true,
			}
			break
		case "password":
			template.Properties[i].Prompt = &promptui.Prompt{
				Label: util.Bold().Sprint(property.Description),
				Mask:  []rune("*")[0],
				Validate: func(input string) error {
					return ValidateRequired(input, currentRequired)
				},
			}
			break
		default:
			template.Properties[i].Prompt = &promptui.Prompt{
				Label: util.Bold().Sprint(property.Description),
				Validate: func(input string) error {
					return ValidateRequired(input, currentRequired)
				},
			}
			break
		}
	}

	return nil
}

func SetPropertiesRequired(props []model.TemplateProperty, propNames []string) []model.TemplateProperty {
	for i, property := range props {
		for _, propName := range propNames {
			if property.Name == propName {
				property.Required = true
				property.Prompt.Validate = func(input string) error {
					return ValidateRequired(input, true)
				}
				props[i] = property
			}
		}
	}
	return props
}

func SetTemplateProps(template *model.Template) error {
	err := PopulatePrompts(template)
	if err != nil {
		return err
	}

	for _, property := range template.Properties {
		if property.Required {
			value, err := util.SimplePrompt(property.Prompt, property.Type == "boolean")
			if err != nil {
				return err
			}

			if property.Required && property.Type == "boolean" && value.(bool) && len(property.Requires) > 0 {
				template.Properties = SetPropertiesRequired(template.Properties, property.Requires)
			}

			if template.PropertyValues == nil {
				template.PropertyValues = make(map[string]interface{})
			}
			template.PropertyValues[property.Name] = value
		}
	}

	return nil
}
