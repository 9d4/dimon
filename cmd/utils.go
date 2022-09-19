package cmd

import (
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

type cobraFunc func(cmd *cobra.Command, args []string)

func wrapCobraFunc(fn func(cmd *cobra.Command, args []string)) cobraFunc {
	return func(cmd *cobra.Command, args []string) {
		bindPFlagsViper(cmd.Flags())
		fn(cmd, args)
	}
}

func newTabWriter() *tabwriter.Writer {
	return tabwriter.NewWriter(os.Stdout, 4, 1, 4, 0x20, 0)
}
