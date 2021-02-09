package util

import (
	"github.com/manifoldco/promptui"
)

func YesNoPrompt(label string) bool {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}

	_, err := prompt.Run()

	if err != nil {
		return false
	}

	return true
}
