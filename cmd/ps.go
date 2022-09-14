package cmd

import (
	"context"

	"github.com/9d4/dimon/client"
	"github.com/spf13/cobra"
)

var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "show running processes",
	Run: wrapCobraFunc(func(cmd *cobra.Command, args []string) {
		cli := client.NewClient()
		cli.ProcessList(context.Background())
	}),
}
