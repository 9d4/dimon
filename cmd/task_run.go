package cmd

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/9d4/dimon/client"
	"github.com/spf13/cobra"
)

var taskRunCmd = &cobra.Command{
	Use:   "run [taskID]",
	Short: "run command by id",
	Run: wrapCobraFunc(func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}

		var taskID int
		if len(args) > 0 {
			i, err := strconv.Atoi(args[0])
			if err != nil {
				cmd.Help()
				return
			}
			taskID = i
		}

		cli := client.NewClient()
		task, err := cli.TaskRun(context.Background(), taskID)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s: %s %s\n", task.Name, task.Command, strings.Join(task.Args, " "))
	}),
}
