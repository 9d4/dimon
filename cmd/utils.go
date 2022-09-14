package cmd

import (
	"github.com/spf13/cobra"
)

type cobraFunc func(cmd *cobra.Command, args []string)

func wrapCobraFunc(fn func(cmd *cobra.Command, args []string)) cobraFunc {
	return func(cmd *cobra.Command, args []string) {
		bindPFlagsViper(cmd.Flags())
		fn(cmd, args)
	}
}
