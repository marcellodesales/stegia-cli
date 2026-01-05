package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"stegia/internal/logger"
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
		"",
		"Override LOG_LEVEL in .env (supported: none, info, debug, error)",
	)

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().Changed("log-level") {
			logger.SetLevelOverride(logLevel)
		}
		return nil
	}
	rootCmd.AddCommand(totvsCmd)
}
