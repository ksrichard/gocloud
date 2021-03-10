package service

import (
	"errors"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/ksrichard/gocloud/util"
	"github.com/spf13/cobra"
	"strings"
)

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
