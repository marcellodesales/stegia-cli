package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	logLevel string
)

var rootCmd = &cobra.Command{
	Use:   "stegia",
	Short: "stegia CLI",
	Long:  "stegia: utilities for integrations and automation",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {

	// Global / persistent flag
	rootCmd.PersistentFlags().StringVarP(
		&logLevel,
		"log-level",
		"l",
		"none",
		"Log level: none, info, debug, error",
	)

	rootCmd.AddCommand(totvsCmd)
}
