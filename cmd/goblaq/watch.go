package main

import (
	"github.com/spf13/cobra"
	"github.com/traksel/goblaq/internal/require"
	"github.com/traksel/goblaq/internal/watch"
)

var watchCmd = &cobra.Command{
	Use:   "watch [NAME] [TARGET]",
	Args:  require.ExactArgs(2),
	Short: "",
	Long:  ``,
	RunE:  runWatch,
}

func runWatch(cmd *cobra.Command, args []string) error {
	name := args[0]
	target := args[1]
	schema, _ := cmd.Flags().GetString("schema")
	path, _ := cmd.Flags().GetString("path")
	return watch.Add(name, target, schema, path)
}

func prepareWatchCmd() {
	watchCmd.PersistentFlags().String("schema", "http", "Set target schema e.g.: --schema https. Default schema: http")
	watchCmd.PersistentFlags().String("path", "/", "Target status path, e.g.: --path \"/app/status\"")
}

func init() {
	prepareWatchCmd()
}
