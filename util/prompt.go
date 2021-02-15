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

func SimplePrompt(prompt *promptui.Prompt, confirmable bool) (interface{}, error) {
	value, err := prompt.Run()

	if err != nil && confirmable {
		return false, nil
	}

	if err != nil && !confirmable {
		return nil, err
	}

	if err == nil && confirmable {
		return true, nil
	}

	return value, nil
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
