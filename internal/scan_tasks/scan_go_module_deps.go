package scan_tasks

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

type ScanGoModuleDepsOptions struct {
	Excludes      []string
	EnabledChecks []Check
	Verbose       bool
}

func ScanGoModuleDeps(opts ScanGoModuleDepsOptions) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	log.Printf("Scanning deps for: %s\n", cwd)
	log.Printf("with checks: %s\n", opts.EnabledChecks)
	log.Printf("excluding: %s\n", opts.Excludes)
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
		if excludeTarget(opts.Excludes, target) {
			if opts.Verbose {
				log.Printf("Skipping excluded target: %s\n", target)
			}
			continue
		}
		_ = ScanGoModule(ScanGoModuleOptions{
			Module:        target,
			EnabledChecks: opts.EnabledChecks,
			Verbose:       opts.Verbose,
		})
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
