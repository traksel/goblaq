package main

import (
	"github.com/spf13/cobra"
)

var webCmd = cobra.Command{
	Use:     "web",
	Aliases: []string{"web-server", "webserver"},
	Short:   `web`,
	Long:    ``,
}

func ExecuteWeb() *cobra.Command {
	return &webCmd
}
