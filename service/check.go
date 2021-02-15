package service

import (
	"errors"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/ksrichard/gocloud/util"
	"os/exec"
)

// TODO: add input options, if they are set they will be used (like automatically install requirements or so...)
func CheckRequirements(requirements []string) error {
	for _, requirement := range requirements {
		err := CheckRequirement(requirement)
		if err != nil {
			return err
		}
	}
	return nil
}

// TODO: add input options, if they are set they will be used (like automatically install requirements or so...)
func CheckRequirement(command string) error {
	err := util.DefaultLoading(func(sp *spinner.Spinner) error {
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
