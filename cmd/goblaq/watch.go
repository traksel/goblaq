package main

import (
	"github.com/spf13/cobra"
)

var watchCmd = cobra.Command{
	Use:   "watch [name] [flags]",
	Short: "",
	Long:  ``,
	RunE:  runWatch,
}

func runWatch(cmd *cobra.Command, args []string) error {
	return watch.appendWatch(args)
}

func prepareWatchCmd() {
	// randomCmd.PersistentFlags().String("term", "", "A search term for a dad joke.")
	watchCmd.PersistentFlags().String("url", "", "Set instance url")
	watchCmd.PersistentFlags().StringSlice("set-service", []string{}, "Set instance url")
}

// func printFlags(cmd *cobra.Command, args []string) {
// 	flag, _ := cmd.Flags().GetString("url")
// 	fmt.Println(flag)
// }

func init() {
	prepareWatchCmd()
}
