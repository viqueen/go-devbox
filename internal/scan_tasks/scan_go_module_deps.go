package scan_tasks

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func ScanGoModuleDeps(enabledChecks []Check) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	log.Printf("Scanning deps for: %s\n", cwd)
	deps, err := listGoModuleDeps()
	if err != nil {
		return err
	}
	for _, dep := range deps {
		parts := strings.Split(dep, " ")
		if len(parts) < 2 {
			continue
		}
		target := parts[0]
		_ = ScanGoModule(target, enabledChecks)
	}
	return nil
}

func listGoModuleDeps() ([]string, error) {
	cmd := exec.Command("go", "list", "-m", "all")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return strings.Split(string(output), "\n"), nil
}
