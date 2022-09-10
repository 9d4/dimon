package cmd

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/9d4/dimon/server"
	"github.com/9d4/dimon/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	v "github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Short: "dimon",
	Long:  "dimon is a simple daemon to run any command as background process",
	Use:   "dimon",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		bindPFlagsViper(cmd.Flags())

		err := storage.Initialize(v.GetString("database"))
		if err != nil {
			log.Fatal(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		server.Start()
	},
}

func init() {
	initConfig()

	rootCmd.PersistentFlags().String("socketpath", path.Join(v.GetString("socketdir"), "sock"), "where the socket will listen on")
	rootCmd.PersistentFlags().StringP("database", "d", "/var/lib/dimon/dimon.db", "database path of dimon")
	rootCmd.PersistentFlags().MarkHidden("socketpath")
	rootCmd.PersistentFlags().MarkHidden("database")
}

func initConfig() {
	v.SetDefault("socketdir", "/var/run/dimon/")
	v.SetDefault("socketmask", 0666)
}

func bindPFlagsViper(flags *pflag.FlagSet) {
	v.BindPFlags(flags)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
