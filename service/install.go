package service

import (
	"bytes"
	"fmt"
	"github.com/ksrichard/gocloud/util"
	"github.com/matishsiao/goInfo"
	"os"
	"os/exec"
	"strings"
)

func GetOS() string {
	gi := goInfo.GetInfo()
	return gi.GoOS
}

func InstallPrompt(toInstall string) error {
	util.BoldRed().Sprintln("'%s' is not installed!", toInstall)
	install := util.YesNoPrompt(util.Bold().Sprintf("Would you like to install '%s'", toInstall))
	if install {
		switch toInstall {
		case "pulumi":
			return InstallPulumi()
		case "kubectl":
			return InstallKubectl()
		}
	}

	return util.BoldError(fmt.Sprintf("'%s' is not installed!", toInstall))
}

// kubectl
func InstallKubectl() error {
	switch GetOS() {
	case "darwin":
		return InstallKubectlDarwin()
	case "linux":
		return InstallKubectlLinux()
		// TODO: add windows installer
	}

	return util.BoldError("Sorry, at the moment we do not support your OS!")
}

func InstallKubectlLinux() error {
	err := util.DefaultLoading(func() error {
		// check curl
		cmdExistsErr := commandExists("curl")
		if cmdExistsErr != nil {
			return cmdExistsErr
		}

		// get version
		err, versionOut := util.RunCmd("curl", "-L", "-s", "https://dl.k8s.io/release/stable.txt")
		if err != nil {
			return err
		}

		// get binary
		err, _ = util.RunCmd("curl", "-LO", "https://dl.k8s.io/release/"+versionOut.String()+"/bin/linux/amd64/kubectl")
		if err != nil {
			return err
		}

		// chmod
		err, _ = util.RunCmd("chmod", "+x", "kubectl")
		if err != nil {
			return err
		}

		// mv
		err, _ = util.RunCmd("mv", "kubectl", "/usr/local/bin/kubectl")
		if err != nil {
			return err
		}

		return err
	}, "Installing Kubectl...", ":robot:")

	if err == nil {
		// print kubectl version
		var versionOut bytes.Buffer
		versionCmd := exec.Command("kubectl", "version", "--client", "--short")
		versionCmd.Stdout = &versionOut
		err = versionCmd.Start()
		if err != nil {
			return err
		}
		err = versionCmd.Wait()
		if err != nil {
			return err
		}
		util.EmojiPrintln(
			":thumbs_up:",
			fmt.Sprint("Kubectl is ready to use with version ", strings.ReplaceAll(versionOut.String(), "\n", ""), "!"),
		)
	}

	return err
}

func InstallKubectlDarwin() error {
	err := util.DefaultLoading(func() error {
		// check curl
		cmdExistsErr := commandExists("curl")
		if cmdExistsErr != nil {
			return cmdExistsErr
		}

		// get version
		err, versionOut := util.RunCmd("curl", "-L", "-s", "https://dl.k8s.io/release/stable.txt")
		if err != nil {
			return err
		}

		// get binary
		err, _ = util.RunCmd("curl", "-LO", "https://dl.k8s.io/release/"+versionOut.String()+"/bin/darwin/amd64/kubectl")
		if err != nil {
			return err
		}

		// chmod
		err, _ = util.RunCmd("chmod", "+x", "kubectl")
		if err != nil {
			return err
		}

		// mv
		err, _ = util.RunCmd("mv", "kubectl", "/usr/local/bin/kubectl")
		if err != nil {
			return err
		}

		return err
	}, "Installing Kubectl...", ":robot:")

	if err == nil {
		// print kubectl version
		var versionOut bytes.Buffer
		versionCmd := exec.Command("kubectl", "version", "--client", "--short")
		versionCmd.Stdout = &versionOut
		err = versionCmd.Start()
		if err != nil {
			return err
		}
		err = versionCmd.Wait()
		if err != nil {
			return err
		}
		util.EmojiPrintln(
			":thumbs_up:",
			fmt.Sprint("Kubectl is ready to use with version ", strings.ReplaceAll(versionOut.String(), "\n", ""), "!"),
		)
	}

	return err
}

// pulumi
func InstallPulumi() error {
	switch GetOS() {
	case "darwin":
		return InstallPulumiUnix()
	case "linux":
		return InstallPulumiUnix()
		// TODO: add windows installer
	}

	return util.BoldError("Sorry, at the moment we do not support your OS!")
}

func PreparePulumiPath() error {
	// set path to include pulumi bin dir
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	err = os.Setenv("PATH", fmt.Sprintf("%s:%s/.pulumi/bin", os.Getenv("PATH"), homeDir))
	if err != nil {
		return err
	}
	return nil
}

func InstallPulumiUnix() error {
	err := PreparePulumiPath()
	if err != nil {
		return err
	}

	err = util.DefaultLoading(func() error {
		// check curl
		cmdExistsErr := commandExists("curl")
		if cmdExistsErr != nil {
			return cmdExistsErr
		}

		// get install sh script
		err, installBashOut := util.RunCmd("curl", "-fsSL", "https://get.pulumi.com")
		if err != nil {
			return err
		}

		// pipe it to sh
		err, _ = util.RunCmdIn(installBashOut, "sh")
		if err != nil {
			return err
		}

		return err
	}, "Installing Pulumi...", ":robot:")

	if err == nil {
		// print pulumi version
		err, versionOut := util.RunCmd("pulumi", "version")
		if err != nil {
			return err
		}

		util.EmojiPrintln(
			":thumbs_up:",
			fmt.Sprint("Pulumi is ready to use with version ", strings.ReplaceAll(versionOut.String(), "\n", ""), "!"),
		)
	}

	return err
}
