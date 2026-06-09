package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dogcli",
	Short: "🐕 A CLI tool to fetch dog images",
	Long:  `DogCLI is a command-line tool that fetches random dog images or images by breed from the Dog API.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(randomCmd)
	rootCmd.AddCommand(breedCmd)
	rootCmd.AddCommand(listCmd)
}
