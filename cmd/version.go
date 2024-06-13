/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"es-client/commons"
)

// VersionCmd represents the version command
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show current version of AList",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(`Built At: %s
Go Version: %s
Author: %s
Commit ID: %s
Version: %s
`,
			commons.BuiltAt, commons.GoVersion, commons.GitAuthor, commons.GitCommit, commons.Version)
		os.Exit(0)
	},
}

func init() {
	RootCmd.AddCommand(VersionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}