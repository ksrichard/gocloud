package service

import (
	"errors"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/ksrichard/gocloud/util"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

func GetOutputVarsFromOtherProjects() (map[string]interface{}, error) {
	result := make(map[string]interface{})

	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if util.IsDir(path) {
				projectYamlFile := fmt.Sprintf("%s/%s", path, ProjectYamlFileName)
				if util.FileExists(projectYamlFile) {
					// get current pulumi user
					currentPulumiUser, err := GetCurrentPulumiUser(path)
					if err != nil {
						return err
					}

					// get project name
					currentPulumiProjectName, err := GetPulumiProjectName(path)
					if err != nil {
						return err
					}

					prefix := fmt.Sprintf("%s/%s", currentPulumiUser, currentPulumiProjectName)
					outputVars, err := GetPulumiProjectVars(path)
					if err != nil {
						return err
					}
					for _, outputVar := range outputVars {
						result[fmt.Sprintf("%s/%s", prefix, outputVar)] = fmt.Sprintf("%s/%s", prefix, outputVar)
					}

				}
			}
			return nil
		})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetOutputVarsWithValuesFromOtherProjects() (map[string]string, error) {
	result := make(map[string]string)

	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if util.IsDir(path) {
				projectYamlFile := fmt.Sprintf("%s/%s", path, ProjectYamlFileName)
				if util.FileExists(projectYamlFile) {
					// get current pulumi user
					currentPulumiUser, err := GetCurrentPulumiUser(path)
					if err != nil {
						return err
					}

					// get project name
					currentPulumiProjectName, err := GetPulumiProjectName(path)
					if err != nil {
						return err
					}

					prefix := fmt.Sprintf("%s/%s", currentPulumiUser, currentPulumiProjectName)
					outputVars, err := GetPulumiProjectVarsWithValues(path)
					if err != nil {
						return err
					}
					for k, v := range outputVars {
						result[fmt.Sprintf("%s/%s", prefix, k)] = v
					}

				}
			}
			return nil
		})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func StartProject(cmd *cobra.Command) error {
	projectDir := cmd.Flag("project-dir").Value.String()
	template, err := GetProjectYaml(projectDir)
	if err != nil {
		return err
	}

	// run start script
	err = util.DefaultLoading(func(sp *spinner.Spinner) error {
		// get install script
		os := GetOS()
		runScript := ""
		for _, script := range template.UpScripts {
			if strings.ToLower(script.OS) == os {
				runScript = script.Script
			}
		}

		if os == "darwin" || os == "linux" {
			// chmod
			chmodErr, _ := util.RunCmdInDir(projectDir, "chmod", "+x", runScript)
			if chmodErr != nil {
				return chmodErr
			}

			sp.Lock()
			fmt.Println()
			// run script
			cmdErr := util.RunCmdInteractiveInDir(projectDir, "/bin/sh", runScript)
			sp.Unlock()
			fmt.Println()

			return cmdErr
		}

		return errors.New("Failed to start project, not supported OS!")
	}, "Starting project...", ":robot:")
	if err != nil {
		return err
	}

	return nil
}

func DestroyProject(cmd *cobra.Command) error {
	projectDir := cmd.Flag("project-dir").Value.String()
	template, err := GetProjectYaml(projectDir)
	if err != nil {
		return err
	}

	// run start script
	err = util.DefaultLoading(func(sp *spinner.Spinner) error {
		// get install script
		os := GetOS()
		runScript := ""
		for _, script := range template.DownScripts {
			if strings.ToLower(script.OS) == os {
				runScript = script.Script
			}
		}

		if os == "darwin" || os == "linux" {
			// chmod
			chmodErr, _ := util.RunCmdInDir(projectDir, "chmod", "+x", runScript)
			if chmodErr != nil {
				return chmodErr
			}

			sp.Lock()
			fmt.Println()
			// run script
			cmdErr := util.RunCmdInteractiveInDir(projectDir, "/bin/sh", runScript)
			sp.Unlock()
			fmt.Println()

			return cmdErr
		}

		return errors.New("Failed to destroy project, not supported OS!")
	}, "Stopping project...", ":robot:")
	if err != nil {
		return err
	}

	return nil
}

func InstallProject(cmd *cobra.Command) error {
	projectDir := cmd.Flag("project-dir").Value.String()
	template, err := GetProjectYaml(projectDir)
	if err != nil {
		return err
	}

	// run start script
	err = util.DefaultLoading(func(sp *spinner.Spinner) error {
		// get install script
		os := GetOS()
		runScript := ""
		for _, script := range template.InstallScripts {
			if strings.ToLower(script.OS) == os {
				runScript = script.Script
			}
		}

		if os == "darwin" || os == "linux" {
			// chmod
			chmodErr, _ := util.RunCmdInDir(projectDir, "chmod", "+x", runScript)
			if chmodErr != nil {
				return chmodErr
			}

			sp.Lock()
			fmt.Println()
			// run script
			cmdErr := util.RunCmdInteractiveInDir(projectDir, "/bin/sh", runScript)
			sp.Unlock()
			fmt.Println()

			return cmdErr
		}

		return errors.New("Failed to install project, not supported OS!")
	}, "Initializing project...", ":robot:")
	if err != nil {
		return err
	}

	return nil
}
