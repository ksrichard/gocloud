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
	"github.com/ksrichard/gocloud/service"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Initialize cloud-native project",
	Long:    `Initialize cloud-native project in a specific directory`,
	PreRunE: service.CheckRequirementsCobra,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: continue flow: select template -> set options for template ()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// TODO: add parameters when flow is done
}
