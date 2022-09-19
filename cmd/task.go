package cmd

import "github.com/spf13/cobra"

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "management task",
}

func init() {
	taskCmd.AddCommand(taskListCmd)
}
