package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "goblaq",
	Short:        "Goblaq is a command-line application monitoring util",
	Long:         ``,
	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(watchCmd)
}

func main() {
	rootCmd.Execute()
}
