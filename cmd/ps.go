package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/9d4/dimon/client"
	"github.com/spf13/cobra"
)

var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "show running processes",
	Run: wrapCobraFunc(func(cmd *cobra.Command, args []string) {
		cli := client.NewClient()
		processes, err := cli.ProcessList(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		tw := tabwriter.NewWriter(os.Stdout, 4, 1, 4, 0x20, 0)
		defer tw.Flush()

		fmt.Fprintln(tw, "ID\tPID\tStatus\tTask\tRun")

		for i, p := range processes {
			if err != nil {
				continue
			}

			var running string = "Running"
			if !p.Status {
				running = "Exited"
			}

			fmt.Fprintf(tw, "%d\t%d\t%s\t%v\t%v\n", i, p.PID, running, p.Task.Name, p.Run)
		}
	}),
}
