package service

import (
	"errors"
	"fmt"
	"github.com/ksrichard/gocloud/util"
	"github.com/spf13/cobra"
	"os/exec"
)

func CheckRequirementsCobra(cmd *cobra.Command, args []string) error {
	// check requirements to run project initialization
	err := CheckRequirements()
	if err != nil {
		return err
	}

	return nil
}

// TODO: add input options, if they are set they will be used (like automatically install requirements or so...)
func CheckRequirements() error {
	err := PreparePulumiPath()
	if err != nil {
		return err
	}

	err = CheckRequirement("pulumi")
	if err != nil {
		return err
	}

	err = CheckRequirement("kubectl")
	if err != nil {
		return err
	}

	return err
}

// TODO: add input options, if they are set they will be used (like automatically install requirements or so...)
func CheckRequirement(command string) error {
	err := util.DefaultLoading(func() error {
		cmdExistsErr := commandExists(command)
		if cmdExistsErr != nil {
			return cmdExistsErr
		}
		return nil
	}, fmt.Sprintf("Checking '%s'...", command), ":face_with_monocle:")

	if err != nil {
		return InstallPrompt(command)
	}

	return err
}

func commandExists(cmd string) error {
	_, err := exec.LookPath(cmd)
	if err != nil {
		return errors.New(fmt.Sprintf("Can't find '%s', please install it!", cmd))
	}
	return err
}
