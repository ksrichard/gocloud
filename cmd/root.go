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
	"github.com/ksrichard/gocloud/util"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var templateDir string
var templateRepo string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gocloud",
	Short: "Kick-start any cloud native project",
	Long:  `Kick-start any cloud native project in your favourite programming language`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// check automatically requirements
	requirementsErr := service.CheckRequirements([]string{"pulumi"})
	if requirementsErr != nil {
		fmt.Println(requirementsErr)
		os.Exit(1)
	}

	// templates git repo
	templateRepo := "https://github.com/ksrichard/gocloud-templates"
	rootCmd.PersistentFlags().StringVarP(&templateRepo, "template-repo", "r", templateRepo,
		`Project template repository, it can be set by setting GOCLOUD_TEMPLATE_REPOSITORY env variable as well`,
	)
	rootCmd.MarkPersistentFlagDirname("template-repo")
	rootCmd.MarkFlagRequired("template-repo")
	viper.AutomaticEnv()
	templateRepoEnv := viper.GetString("gocloud_template_repository")
	if templateRepoEnv != "" {
		templateRepo = templateRepoEnv
	}

	// template download dir
	homeDir, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	templateDir := homeDir + "/gocloud-templates"
	rootCmd.PersistentFlags().StringVarP(&templateDir, "template-dir", "t", templateDir,
		`Template directory to download and store project templates,
it can be set by setting GOCLOUD_TEMPLATE_DIR env variable as well`,
	)
	rootCmd.MarkPersistentFlagDirname("template-dir")
	rootCmd.MarkFlagRequired("template-dir")
	viper.AutomaticEnv()
	templateDirEnv := viper.GetString("gocloud_template_dir")
	if templateDirEnv != "" {
		templateDir = templateDirEnv
	}

	// create/check templates dir
	initErr := util.InitCliHomeDir(templateDir)
	if initErr != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cloneErr := util.CloneAndPullTemplatesRepo(templateRepo, templateDir)
	if cloneErr != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
