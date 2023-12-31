package require

import (
	"fmt"

	"github.com/spf13/cobra"
)

func ExactArgs(n int) cobra.PositionalArgs {
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
