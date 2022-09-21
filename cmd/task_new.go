package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/9d4/dimon/client"
	"github.com/spf13/cobra"
)

var taskNewCmd = &cobra.Command{
	Use:                   "new [name] [command] -- [args...]",
	Short:                 "create new task",
	DisableFlagsInUseLine: true,
	Run: wrapCobraFunc(func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Help()
			return
		}

		argName := args[0]
		argCmd := args[1]
		argArgs := []string{}

		if len(args) > 2 {
			argArgs = args[2:]
		}

		cli := client.NewClient()
		task, err := cli.TaskNew(context.Background(), argName, argCmd, argArgs...)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Name   : %s \n", task.Name)
		fmt.Printf("Command: %s \n", task.Command)
		fmt.Printf("Args   : %s \n", strings.Join(task.Args, " "))
	}),
}
