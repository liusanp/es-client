package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "es-client",
	Short: "An easy-to-use Elasticsearch query web page.",
	Run: func(cmd *cobra.Command, args []string) {
		ServerCmd.Run(cmd, args)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
