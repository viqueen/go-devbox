package scan_tasks

import (
	gotasks "github.com/viqueen/go-devbox/internal/go_tasks"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type ScanGoModuleOptions struct {
	Module        string
	EnabledChecks []Check
	Verbose       bool
}

func ScanGoModule(opts ScanGoModuleOptions) error {
	if opts.Verbose {
		log.Printf("Scanning module: %s\n", opts.Module)
	}
	info, err := gotasks.List(opts.Module)
	if err != nil {
		return err
	}
	goFiles, err := listGoFiles(info.Dir)
	if err != nil {
		return err
	}
	for _, file := range goFiles {
		scanErr := ScanGoFile(file, opts.EnabledChecks)
		if scanErr != nil && opts.Verbose {
			log.Printf("Error scanning file %s: %v\n", file, scanErr)
		}
	}
	return nil
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
