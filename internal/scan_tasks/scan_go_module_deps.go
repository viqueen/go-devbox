package scan_tasks

import (
	"log"
	"os"
)

func ScanGoModuleDeps() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	log.Printf("Scanning deps for: %s\n", cwd)
	return nil
}
