package scan_tasks

import (
	gotasks "github.com/viqueen/go-devbox/internal/go_tasks"
	"log"
	"os"
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
	err = gotasks.ModTidy()
	if err != nil {
		return err
	}
	deps, err := gotasks.ListAll()
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
