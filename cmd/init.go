/*
Copyright © 2021 Richard Klavora <klavorasr@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/ksrichard/gocloud/model"
	"github.com/ksrichard/gocloud/service"
	"github.com/ksrichard/gocloud/util"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize cloud-native project",
	Long:  `Initialize cloud-native project in a specific directory`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: continue flow: select template -> set options for template ()
		templateDir := cmd.Flag("template-dir").Value.String()

		// get template project yaml
		templateProject, err := service.GetTemplateProjectYaml(templateDir)
		if err != nil {
			return err
		}

		// select root project
		var templateProjectItemsMapping = make(map[string]interface{})
		for _, project := range templateProject.Projects {
			templateProjectItemsMapping[project.Title] = project
		}
		selectedProject, err := util.Select("Please select a project", templateProjectItemsMapping)
		if err != nil {
			return err
		}
		selectedProjectFolder := fmt.Sprintf("%s/%s", templateDir, selectedProject.(model.TemplateProject).Folder)

		// select template
		dirs, err := service.GetAllTemplates(selectedProjectFolder)
		if err != nil {
			return err
		}

		fmt.Printf("%v", dirs)

		// TODO: continue flow

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// TODO: add parameters when flow is done
}
