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
	"github.com/ksrichard/gocloud/service"
	"os"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Setup and install project dependencies and switch to correct context",
	Long:  `Setup and install project dependencies and switch to correct context`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return service.InstallProject(cmd)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	installCmd.Flags().StringP("project-dir", "p", currentDir, "Project directory to use")
}
