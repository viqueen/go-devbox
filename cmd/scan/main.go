package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/viqueen/go-devbox/internal/scan_tasks"
	"os"
)

var rootCmd = &cobra.Command{}

var scanModCmd = &cobra.Command{
	Use:   "mod",
	Args:  cobra.MinimumNArgs(1),
	Short: "scan a go module",
	RunE: func(cmd *cobra.Command, args []string) error {
		target := args[0]
		return scan_tasks.ScanGoModule(target)
	},
}

var scanDepsCmd = &cobra.Command{
	Use:   "deps",
	Short: "scan dependencies of a go module",
	RunE: func(cmd *cobra.Command, args []string) error {
		return scan_tasks.ScanGoModuleDeps()
	},
}

func init() {
	rootCmd.AddCommand(scanModCmd)
	rootCmd.AddCommand(scanDepsCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
