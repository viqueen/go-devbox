package scan_tasks

import "strings"

func excludeTarget(excludes []string, target string) bool {
	for _, exclude := range excludes {
		if exclude != "" && strings.Contains(target, exclude) {
			return true
		}
	}
	return false
}
