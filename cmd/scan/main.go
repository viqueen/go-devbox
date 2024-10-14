package main

import (
	"fmt"
	"github.com/spf13/cobra"
	scantasks "github.com/viqueen/go-devbox/internal/scan_tasks"
	"os"
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
		})
	},
}

var scanDepsCmd = &cobra.Command{
	Use:   "deps",
	Short: "scan dependencies of a go module",
	RunE: func(cmd *cobra.Command, args []string) error {
		checks := cmd.Flag("checks").Value.String()
		verbose, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			return err
		}
		return scantasks.ScanGoModuleDeps(scantasks.ScanGoModuleDepsOptions{
			EnabledChecks: scantasks.ParseChecks(checks),
			Verbose:       verbose,
		})
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("checks", "c", "", "comma separated list of checks to run")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.AddCommand(scanModCmd)
	rootCmd.AddCommand(scanDepsCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
