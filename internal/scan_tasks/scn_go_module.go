package scan_tasks

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ScanGoModule(module string, enabledChecks []Check) error {
	log.Printf("Scanning module: %s\n", module)
	info, err := getModInfo(module)
	if err != nil {
		return err
	}
	goFiles, err := listGoFiles(info.Dir)
	if err != nil {
		return err
	}
	for _, file := range goFiles {
		scanErr := ScanGoFile(file, enabledChecks)
		if scanErr != nil {
			log.Printf("Error scanning file %s: %v\n", file, scanErr)
		}
	}
	return nil
}

type modInfo struct {
	Dir string `json:"Dir"`
}

func getModInfo(module string) (*modInfo, error) {
	cmd := exec.Command("go", "list", "-json", module)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var info modInfo
	if jsonErr := json.Unmarshal(output, &info); jsonErr != nil {
		return nil, jsonErr
	}

	return &info, nil
}

func listGoFiles(moduleDir string) ([]string, error) {
	var goFiles []string
	err := filepath.Walk(moduleDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go") {
			goFiles = append(goFiles, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return goFiles, nil
}
