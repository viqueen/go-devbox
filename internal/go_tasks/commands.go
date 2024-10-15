package go_tasks

import (
	"encoding/json"
	"os/exec"
	"strings"
)

func ModTidy() error {
	cmd := exec.Command("go", "mod", "tidy")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func Get(module string) error {
	cmd := exec.Command("go", "get", module)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

type ModInfo struct {
	Dir string `json:"Dir"`
}

func List(module string) (*ModInfo, error) {
	cmd := exec.Command("go", "list", "-json", module)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var info ModInfo
	if jsonErr := json.Unmarshal(output, &info); jsonErr != nil {
		return nil, jsonErr
	}

	return &info, nil
}

func ListAll() ([]string, error) {
	cmd := exec.Command("go", "list", "-m", "all")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return strings.Split(string(output), "\n"), nil
}
