package main

import (
	"fmt"
	"github.com/spf13/cobra"
	scantasks "github.com/viqueen/go-devbox/internal/scan_tasks"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{}

var scanModCmd = &cobra.Command{
	Use:   "mod",
	Args:  cobra.MinimumNArgs(1),
	Short: "scan a go module",
	RunE: func(cmd *cobra.Command, args []string) error {
		checks := cmd.Flag("checks").Value.String()
		verbose, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			return err
		}
		target := args[0]
		return scantasks.ScanGoModule(scantasks.ScanGoModuleOptions{
			Module:        target,
			EnabledChecks: scantasks.ParseChecks(checks),
			Verbose:       verbose,
			WithGet:       true,
		})
	},
}

var scanDepsCmd = &cobra.Command{
	Use:   "deps",
	Short: "scan dependencies of a go module",
	RunE: func(cmd *cobra.Command, args []string) error {
		checks := cmd.Flag("checks").Value.String()
		excludes := cmd.Flag("exclude").Value.String()
		verbose, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			return err
		}
		tidy, err := cmd.Flags().GetBool("tidy")
		if err != nil {
			return err
		}
		return scantasks.ScanGoModuleDeps(scantasks.ScanGoModuleDepsOptions{
			EnabledChecks: scantasks.ParseChecks(checks),
			Excludes:      strings.Split(excludes, ","),
			Verbose:       verbose,
			WithTidy:      tidy,
		})
	},
}

var scanGitHubCmd = &cobra.Command{
	Use:   "github",
	Short: "scan a github go module",
	RunE: func(cmd *cobra.Command, args []string) error {
		checks := cmd.Flag("checks").Value.String()
		verbose, err := cmd.Flags().GetBool("verbose")
		excludes := cmd.Flag("exclude").Value.String()
		if err != nil {
			return err
		}
		return scantasks.ScanGithub(scantasks.ScanGithubOptions{
			EnabledChecks: scantasks.ParseChecks(checks),
			Excludes:      strings.Split(excludes, ","),
			Verbose:       verbose,
		})
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("checks", "c", "", "comma separated list of checks to run")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringP("exclude", "e", "", "comma separated list of module providers to exclude")
	scanDepsCmd.Flags().BoolP("tidy", "t", false, "run go mod tidy before scanning")

	rootCmd.AddCommand(scanModCmd)
	rootCmd.AddCommand(scanDepsCmd)
	rootCmd.AddCommand(scanGitHubCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
