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

func Select(label string, items map[string]interface{}) (interface{}, error) {
	var promptItems []string
	for k, _ := range items {
		promptItems = append(promptItems, k)
	}

	prompt := promptui.Select{
		Label: Bold().Sprint(label),
		Items: promptItems,
	}

	_, promptResult, err := prompt.Run()

	if err != nil {
		return "", err
	}

	var result interface{}
	for k, v := range items {
		if k == promptResult {
			result = v
		}
	}

	return result, nil
}
