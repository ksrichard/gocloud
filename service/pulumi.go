package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ksrichard/gocloud/model"
	"github.com/ksrichard/gocloud/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

func GetCurrentPulumiUser(projectDir string) (string, error) {
	err, whoAmiOut := util.RunCmdInDir(projectDir, "pulumi", "whoami")
	if err != nil {
		return "", err
	}
	currentPulumiUser := strings.ReplaceAll(whoAmiOut.String(), "\n", "")
	if strings.TrimSpace(currentPulumiUser) == "" {
		return "", errors.New(fmt.Sprintf("No Pulumi user has been logged in in project '%s'!", projectDir))
	}
	return currentPulumiUser, nil
}

func GetPulumiProjectName(projectDir string) (string, error) {
	pulumiYamlFile := fmt.Sprintf("%s/%s", projectDir, "Pulumi.yaml")
	if !util.FileExists(pulumiYamlFile) {
		return "", errors.New(fmt.Sprintf("'%s' not found!", pulumiYamlFile))
	}
	var pulumiYaml model.PulumiYaml
	yamlFile, err := ioutil.ReadFile(pulumiYamlFile)
	if err != nil {
		return "", err
	}
	err = yaml.Unmarshal(yamlFile, &pulumiYaml)
	if err != nil {
		return "", err
	}
	return pulumiYaml.Name, nil
}

func GetPulumiProjectVars(projectDir string) ([]string, error) {
	err, jsonOut := util.RunCmdInDir(projectDir, "pulumi", "stack", "output", "--show-secrets", "--json")
	if err != nil {
		return nil, err
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(jsonOut.Bytes(), &jsonMap)
	if err != nil {
		return nil, err
	}

	var result []string
	for k, _ := range jsonMap {
		result = append(result, k)
	}

	return result, nil
}

func GetPulumiProjectVarsWithValues(projectDir string) (map[string]string, error) {
	err, jsonOut := util.RunCmdInDir(projectDir, "pulumi", "stack", "output", "--show-secrets", "--json")
	if err != nil {
		return nil, err
	}

	jsonMap := make(map[string]string)
	err = json.Unmarshal(jsonOut.Bytes(), &jsonMap)
	if err != nil {
		return nil, err
	}

	return jsonMap, nil
}
