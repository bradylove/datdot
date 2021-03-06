package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "datdot",
	Short: "datdot is a simple dotfile manager",
	Long: `datdot is a simple dotfile manager that uses a git repo for syncing
your dotfiles.`,
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigName("datdot")     // name of config file (without extension)
	viper.AddConfigPath("$HOME/.dot") // adding home directory as first search path

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		var hasHelp bool
		if len(os.Args) > 1 && os.Args[1] == "help" {
			hasHelp = true
		}

		if !hasHelp && os.Args[1] != "init" {
			fmt.Println("datdot is not initialized")
			os.Exit(1)
		}
	}
}
