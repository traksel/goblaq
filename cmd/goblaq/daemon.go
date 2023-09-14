package main

import (
	"github.com/spf13/cobra"
	"github.com/traksel/goblaq/internal/daemon"
	"github.com/traksel/goblaq/internal/require"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon [MODE]",
	Args:  require.ExactArgs(1),
	Short: "",
	Long:  ``,
	Run:   runDaemon,
}

func runDaemon(cmd *cobra.Command, args []string) {
	daemon.Run(args[0])
}
