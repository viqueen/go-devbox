package scan_tasks

import (
	"github.com/viqueen/go-devbox/internal/github"
	"log"
	"strings"
)

type ScanGithubOptions struct {
	Excludes      []string
	EnabledChecks []Check
	Verbose       bool
}

func ScanGithub(opts ScanGithubOptions) error {
	githubClient := github.NewClient()
	repos, err := githubClient.SearchRepositories("")
	if err != nil {
		return err
	}
	for _, repo := range repos {
		module := strings.TrimPrefix(repo.HtmlUrl, "https://")
		if excludeTarget(opts.Excludes, module) {
			if opts.Verbose {
				log.Printf("Skipping excluded target: %s\n", repo.HtmlUrl)
			}
		}
		_ = ScanGoModule(ScanGoModuleOptions{
			Module:        module,
			EnabledChecks: opts.EnabledChecks,
			Verbose:       opts.Verbose,
		})
	}
	return nil
}
