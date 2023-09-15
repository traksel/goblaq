package main

import (
	"github.com/spf13/cobra"
	"github.com/traksel/goblaq/internal/remove"
	"github.com/traksel/goblaq/internal/require"
)

var removeCmd = &cobra.Command{
	Use:   "remove [NAME]",
	Args:  require.ExactArgs(1),
	Short: "",
	Long:  ``,
	RunE:  runRemove,
}

func runRemove(cmd *cobra.Command, args []string) error {
	name := args[0]
	return remove.Remove(name)
}
