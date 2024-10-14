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
		target := args[0]
		return scantasks.ScanGoModule(target, scantasks.ParseChecks(checks))
	},
}

var scanDepsCmd = &cobra.Command{
	Use:   "deps",
	Short: "scan dependencies of a go module",
	RunE: func(cmd *cobra.Command, args []string) error {
		checks := cmd.Flag("checks").Value.String()
		return scantasks.ScanGoModuleDeps(scantasks.ParseChecks(checks))
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("checks", "c", "", "comma separated list of checks to run")
	rootCmd.AddCommand(scanModCmd)
	rootCmd.AddCommand(scanDepsCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
