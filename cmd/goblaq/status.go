package main

import (
	"github.com/spf13/cobra"
	"github.com/traksel/goblaq/internal/require"
	"github.com/traksel/goblaq/internal/status"
)

var statusCmd = &cobra.Command{
	Use:   "status [NAME]",
	Args:  require.ExactArgs(1),
	Short: "",
	Long:  ``,
	Run:   runStatus,
}

func runStatus(cmd *cobra.Command, args []string) {
	s := status.Status{}
	arg := args[0]
	s.Get(arg)
}
