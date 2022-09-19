package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/9d4/dimon/client"
	"github.com/spf13/cobra"
)

var taskListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "show task list",
	Run: wrapCobraFunc(func(cmd *cobra.Command, args []string) {
		cli := client.NewClient()
		tasks, err := cli.TaskList(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		tw := newTabWriter()
		defer tw.Flush()

		fmt.Fprintln(tw, "ID\tName\tCommand")
		for _, t := range tasks {
			fmt.Fprintf(tw, "%d\t%s\t%s\n", t.ID, t.Name, t.CommandArgs)
		}
	}),
}
