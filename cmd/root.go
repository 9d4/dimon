package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/9d4/dimon/server"
	"github.com/spf13/cobra"
	v "github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Short: "dimon",
	Long:  "dimon is a simple daemon to run any command as background process",
	Use:   "dimon",
	Run: func(cmd *cobra.Command, args []string) {
		server.Start()
	},
}

func init() {
	initConfig()
}

func initConfig() {
	v.SetDefault("socketdir", "/var/run/dimon/")
	v.SetDefault("socketpath", path.Join(v.GetString("socketdir"), "sock"))
	v.SetDefault("socketmask", 0666)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
