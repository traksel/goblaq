package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/traksel/goblaq/internal/watch"
)

var watchCmd = &cobra.Command{
	Use:               "watch [NAME] [TARGET]",
	Args:              exactArgs(2),
	Short:             "",
	Long:              ``,
	RunE:              runWatch,
	ValidArgsFunction: noCompletions,
}

func noCompletions(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return nil, cobra.ShellCompDirectiveNoFileComp
}

func exactArgs(n int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) != n {
			return fmt.Errorf(
				"%q requires %d %s\n\nUsage:  %s",
				cmd.CommandPath(),
				n,
				"argument(s)",
				cmd.UseLine(),
			)
		}
		return nil
	}
}

func runWatch(cmd *cobra.Command, args []string) error {
	name := args[0]
	target := args[1]
	schema, _ := cmd.Flags().GetString("schema")
	targetService, _ := cmd.Flags().GetStringSlice("target-service")
	return watch.AppendWatch(name, target, schema, targetService)
}

func prepareWatchCmd() {
	watchCmd.PersistentFlags().String("schema", "", "Set target schema e.g.: --schema https. Default schema: http")
	watchCmd.PersistentFlags().StringSlice("target-service", []string{}, "List of target hosted services, e.g.: --target-service \"/app1, /app2\"")
}

func init() {
	prepareWatchCmd()
}
