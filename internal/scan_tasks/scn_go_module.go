package scan_tasks

import "log"

func ScanGoModule(module string) error {
	log.Printf("Scanning module: %s\n", module)
	return nil
}
