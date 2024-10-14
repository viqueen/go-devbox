package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{}

var scanModCmd = &cobra.Command{
	Use: "mod",
}

var scanDepsCmd = &cobra.Command{
	Use: "deps",
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
