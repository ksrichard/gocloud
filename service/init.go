package service

import (
	"errors"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/ksrichard/gocloud/model"
	"github.com/ksrichard/gocloud/util"
	gogen "github.com/ksrichard/gogen/service"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func InitProject(cmd *cobra.Command) error {
	templateDir := cmd.Flag("template-dir").Value.String()
	targetDir := cmd.Flag("target-dir").Value.String()
	pulumiLocalLogin, _ := cmd.Flags().GetBool("pulumi-local")

	// get template project yaml
	templateProject, err := GetTemplateProjectYaml(templateDir)
	if err != nil {
		return err
	}

	// select root project
	var templateProjectItemsMapping = make(map[string]interface{})
	for _, project := range templateProject.Projects {
		templateProjectItemsMapping[util.Bold().Sprint(project.Title)] = project
	}
	selectedProject, err := util.Select("Please select a project", templateProjectItemsMapping)
	if err != nil {
		return err
	}
	selectedProjectFolder := fmt.Sprintf("%s/%s", templateDir, selectedProject.(model.TemplateProject).Folder)

	// select template
	dirs, err := GetAllTemplates(selectedProjectFolder)
	if err != nil {
		return err
	}
	var templateMapping = make(map[string]interface{})
	for _, dir := range dirs {
		key := fmt.Sprintf("%s - %s", util.Bold().Sprint(dir.Name), dir.Description)
		templateMapping[key] = dir
	}
	selectedTemplate, err := util.Select("Please select a project template", templateMapping)
	if err != nil {
		return err
	}
	template := selectedTemplate.(model.Template)

	// check template additional requirements
	err = CheckRequirements(template.Requirements)
	if err != nil {
		return err
	}

	// set properties of template
	err = SetTemplateProps(&template)
	if err != nil {
		return err
	}

	// pass local login option to template generation
	template.PropertyValues["pulumi_local_login"] = pulumiLocalLogin

	// generate project files
	err = util.DefaultLoading(func(sp *spinner.Spinner) error {
		err := gogen.Generate(template.Folder, targetDir, "folder", template.PropertyValues)
		if err != nil {
			return err
		}

		// delete project template yaml
		projectTemplateYaml, err := GetProjectTemplateYaml(targetDir)
		if err != nil {
			return err
		}
		err = os.Remove(fmt.Sprintf("%s/%s", targetDir, ProjectTemplateYamlFileName))
		if err != nil {
			return err
		}

		// create project specific yaml
		projectYaml := &model.ProjectConfig{
			InstallScripts: projectTemplateYaml.InstallScripts,
			UpScripts:      projectTemplateYaml.UpScripts,
			DownScripts:    projectTemplateYaml.DownScripts,
		}
		err = WriteProjectYaml(targetDir, projectYaml)

		return err
	}, "Generating new project...", ":robot:")
	if err != nil {
		return err
	}

	// run install script
	err = util.DefaultLoading(func(sp *spinner.Spinner) error {
		// get install script
		os := GetOS()
		installScript := ""
		for _, script := range template.InstallScripts {
			if strings.ToLower(script.OS) == os {
				installScript = script.Script
			}
		}

		if os == "darwin" || os == "linux" {
			// chmod
			chmodErr, _ := util.RunCmdInDir(targetDir, "chmod", "+x", installScript)
			if chmodErr != nil {
				return chmodErr
			}

			sp.Lock()
			fmt.Println()
			// run install script
			cmdErr := util.RunCmdInteractiveInDir(targetDir, "/bin/sh", installScript)
			sp.Unlock()
			fmt.Println()

			return cmdErr
		}

		return errors.New("Failed to initialize project, not supported OS!")
	}, "Initializing project...", ":robot:")
	if err != nil {
		return err
	}

	return nil
}
