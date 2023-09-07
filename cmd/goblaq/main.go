package main

import "github.com/spf13/cobra"

var rootCmd = cobra.Command{
	Use:   "goblaq",
	Short: "Goblaq is a command-line application monitoring util",
	Long:  ``,
}

func main() {
	rootCmd.AddCommand(&webCmd)
	rootCmd.Execute()
}
