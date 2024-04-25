package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "task-manager",
	Short: "A simple task manager",
}